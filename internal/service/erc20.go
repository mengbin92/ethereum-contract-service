// Package service provides business logic services for ERC20 token interactions.
package service

import (
	"context"
	"math/big"

	pb "eth-contract-service/api/erc20/v1"
	"eth-contract-service/internal/contract"
	"eth-contract-service/internal/errors"
	"eth-contract-service/internal/validator"
	"eth-contract-service/provider/contract/erc20"
	"eth-contract-service/provider/eth"
	"eth-contract-service/provider/keystore"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kratos/kratos/v2/log"
)

// ERC20Service implements the ERC20 API service.
// It provides methods for interacting with ERC20 tokens.
type ERC20Service struct {
	pb.UnimplementedERC20Server
	logger         *log.Helper
	contractClient *contract.Client
}

// NewERC20Service creates a new instance of ERC20Service.
//
// Parameters:
//   - logger: Logger instance for service logging
//
// Returns:
//   - *ERC20Service: A new service instance
func NewERC20Service(logger log.Logger) *ERC20Service {
	return &ERC20Service{
		logger:         log.NewHelper(logger),
		contractClient: contract.NewClient(logger),
	}
}

// GetERC20Balance returns the ERC20 token balance of the specified address.
func (s *ERC20Service) GetERC20Balance(ctx context.Context, req *pb.GetERC20BalanceRequest) (*pb.GetERC20BalanceResponse, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	ownerAddr, err := validator.ValidateAddress(req.OwnerAddress, "owner_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Get contract instance (supports both standard and ownable)
	contractType := req.GetContractType()
	if contractType == "" {
		contractType = "standard"
	}
	token, err := s.getERC20Contract(contractAddr, contractType)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 contract: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get balance
	balance, err := token.BalanceOf(nil, ownerAddr)
	if err != nil {
		s.logger.Errorf("failed to get balance: contract=%s, owner=%s, error=%v", contractAddr.Hex(), ownerAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get balance"))
	}

	// Get decimals for display
	decimals, err := token.Decimals(nil)
	if err != nil {
		s.logger.Warnf("failed to get decimals, using 18 as default: contract=%s, error=%v", contractAddr.Hex(), err)
		decimals = 18
	}

	s.logger.Infof("balance queried: contract=%s, owner=%s, balance=%s", contractAddr.Hex(), ownerAddr.Hex(), balance.String())

	return &pb.GetERC20BalanceResponse{
		Balance:         balance.String(),
		ContractAddress: req.ContractAddress,
		OwnerAddress:    req.OwnerAddress,
		Decimals:        uint32(decimals),
	}, nil
}

// GetERC20Info returns ERC20 token information.
func (s *ERC20Service) GetERC20Info(ctx context.Context, req *pb.GetERC20InfoRequest) (*pb.GetERC20InfoResponse, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate address
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Get contract instance (supports both standard and ownable)
	contractType := req.GetContractType()
	if contractType == "" {
		contractType = "standard"
	}
	token, err := s.getERC20Contract(contractAddr, contractType)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 contract: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get token info
	name, err := token.Name(nil)
	if err != nil {
		s.logger.Errorf("failed to get token name: contract=%s, error=%v", contractAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get name"))
	}

	symbol, err := token.Symbol(nil)
	if err != nil {
		s.logger.Errorf("failed to get token symbol: contract=%s, error=%v", contractAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get symbol"))
	}

	decimals, err := token.Decimals(nil)
	if err != nil {
		s.logger.Errorf("failed to get token decimals: contract=%s, error=%v", contractAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get decimals"))
	}

	totalSupply, err := token.TotalSupply(nil)
	if err != nil {
		s.logger.Errorf("failed to get token total supply: contract=%s, error=%v", contractAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get total supply"))
	}

	s.logger.Infof("token info queried: contract=%s, name=%s, symbol=%s", contractAddr.Hex(), name, symbol)

	return &pb.GetERC20InfoResponse{
		Name:            name,
		Symbol:          symbol,
		Decimals:        uint32(decimals),
		TotalSupply:     totalSupply.String(),
		ContractAddress: req.ContractAddress,
	}, nil
}

// TransferERC20 transfers ERC20 tokens.
func (s *ERC20Service) TransferERC20(ctx context.Context, req *pb.TransferERC20Request) (*pb.TransferERC20Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	toAddr, err := validator.ValidateAddress(req.ToAddress, "to_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate amount
	amount, err := validator.ValidateAmount(req.Amount, "amount")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Get sender address from private key
	fromAddr, err := s.contractClient.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC20Token contract instance
	token, err := s.contractClient.GetERC20Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Transfer tokens
	tx, err := token.Transfer(auth, toAddr, amount)
	if err != nil {
		s.logger.Errorf("failed to transfer tokens: contract=%s, from=%s, to=%s, amount=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to transfer tokens"))
	}

	txHash := tx.Hash()
	s.logger.Infof("transfer initiated: contract=%s, from=%s, to=%s, amount=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), amount.String(), txHash.Hex())

	return &pb.TransferERC20Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     fromAddr.Hex(),
		ToAddress:       req.ToAddress,
		Amount:          req.Amount,
	}, nil
}

// ApproveERC20 approves the spender to spend ERC20 tokens.
func (s *ERC20Service) ApproveERC20(ctx context.Context, req *pb.ApproveERC20Request) (*pb.ApproveERC20Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	spenderAddr, err := validator.ValidateAddress(req.SpenderAddress, "spender_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate amount
	amount, err := validator.ValidateAmount(req.Amount, "amount")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Get owner address from private key
	ownerAddr, err := s.contractClient.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC20Token contract instance
	token, err := s.contractClient.GetERC20Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Approve tokens
	tx, err := token.Approve(auth, spenderAddr, amount)
	if err != nil {
		s.logger.Errorf("failed to approve tokens: contract=%s, owner=%s, spender=%s, amount=%s, error=%v",
			contractAddr.Hex(), ownerAddr.Hex(), spenderAddr.Hex(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to approve tokens"))
	}

	txHash := tx.Hash()
	s.logger.Infof("approval initiated: contract=%s, owner=%s, spender=%s, amount=%s, tx=%s",
		contractAddr.Hex(), ownerAddr.Hex(), spenderAddr.Hex(), amount.String(), txHash.Hex())

	return &pb.ApproveERC20Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		OwnerAddress:    ownerAddr.Hex(),
		SpenderAddress:  req.SpenderAddress,
		Amount:          req.Amount,
	}, nil
}

// GetERC20Allowance returns the amount of tokens that the spender is allowed to spend.
func (s *ERC20Service) GetERC20Allowance(ctx context.Context, req *pb.GetERC20AllowanceRequest) (*pb.GetERC20AllowanceResponse, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	ownerAddr, err := validator.ValidateAddress(req.OwnerAddress, "owner_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	spenderAddr, err := validator.ValidateAddress(req.SpenderAddress, "spender_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC20Token contract instance
	token, err := s.contractClient.GetERC20Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get allowance
	allowance, err := token.Allowance(nil, ownerAddr, spenderAddr)
	if err != nil {
		s.logger.Errorf("failed to get allowance: contract=%s, owner=%s, spender=%s, error=%v",
			contractAddr.Hex(), ownerAddr.Hex(), spenderAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get allowance"))
	}

	s.logger.Infof("allowance queried: contract=%s, owner=%s, spender=%s, allowance=%s",
		contractAddr.Hex(), ownerAddr.Hex(), spenderAddr.Hex(), allowance.String())

	return &pb.GetERC20AllowanceResponse{
		Allowance:       allowance.String(),
		ContractAddress: req.ContractAddress,
		OwnerAddress:    req.OwnerAddress,
		SpenderAddress:  req.SpenderAddress,
	}, nil
}

// TransferFromERC20 transfers ERC20 tokens from one address to another (requires approval).
func (s *ERC20Service) TransferFromERC20(ctx context.Context, req *pb.TransferFromERC20Request) (*pb.TransferFromERC20Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	fromAddr, err := validator.ValidateAddress(req.FromAddress, "from_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	toAddr, err := validator.ValidateAddress(req.ToAddress, "to_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate amount
	amount, err := validator.ValidateAmount(req.Amount, "amount")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC20Token contract instance
	token, err := s.contractClient.GetERC20Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Transfer from
	tx, err := token.TransferFrom(auth, fromAddr, toAddr, amount)
	if err != nil {
		s.logger.Errorf("failed to transfer from: contract=%s, from=%s, to=%s, amount=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to transfer from"))
	}

	txHash := tx.Hash()
	s.logger.Infof("transfer from initiated: contract=%s, from=%s, to=%s, amount=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), amount.String(), txHash.Hex())

	return &pb.TransferFromERC20Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     req.FromAddress,
		ToAddress:       req.ToAddress,
		Amount:          req.Amount,
	}, nil
}

// MintERC20 mints new ERC20 tokens.
func (s *ERC20Service) MintERC20(ctx context.Context, req *pb.MintERC20Request) (*pb.MintERC20Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	toAddr, err := validator.ValidateAddress(req.ToAddress, "to_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate amount
	amount, err := validator.ValidateAmount(req.Amount, "amount")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC20Token contract instance
	token, err := s.contractClient.GetERC20Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Mint tokens
	tx, err := token.Mint(auth, toAddr, amount)
	if err != nil {
		s.logger.Errorf("failed to mint tokens: contract=%s, to=%s, amount=%s, error=%v",
			contractAddr.Hex(), toAddr.Hex(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to mint tokens"))
	}

	txHash := tx.Hash()
	s.logger.Infof("mint initiated: contract=%s, to=%s, amount=%s, tx=%s",
		contractAddr.Hex(), toAddr.Hex(), amount.String(), txHash.Hex())

	return &pb.MintERC20Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		ToAddress:       req.ToAddress,
		Amount:          req.Amount,
	}, nil
}

// BurnERC20 burns ERC20 tokens from the caller's balance.
func (s *ERC20Service) BurnERC20(ctx context.Context, req *pb.BurnERC20Request) (*pb.BurnERC20Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate address
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate amount
	amount, err := validator.ValidateAmount(req.Amount, "amount")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Get sender address from private key
	fromAddr, err := s.contractClient.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC20Token contract instance
	token, err := s.contractClient.GetERC20Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Burn tokens
	tx, err := token.Burn(auth, amount)
	if err != nil {
		s.logger.Errorf("failed to burn tokens: contract=%s, from=%s, amount=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to burn tokens"))
	}

	txHash := tx.Hash()
	s.logger.Infof("burn initiated: contract=%s, from=%s, amount=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), amount.String(), txHash.Hex())

	return &pb.BurnERC20Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     fromAddr.Hex(),
		Amount:          req.Amount,
	}, nil
}

// BurnFromERC20 burns ERC20 tokens from a specified address (requires approval).
func (s *ERC20Service) BurnFromERC20(ctx context.Context, req *pb.BurnFromERC20Request) (*pb.BurnFromERC20Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	fromAddr, err := validator.ValidateAddress(req.FromAddress, "from_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate amount
	amount, err := validator.ValidateAmount(req.Amount, "amount")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC20Token contract instance
	token, err := s.contractClient.GetERC20Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC20 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Burn from
	tx, err := token.BurnFrom(auth, fromAddr, amount)
	if err != nil {
		s.logger.Errorf("failed to burn from: contract=%s, from=%s, amount=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to burn from"))
	}

	txHash := tx.Hash()
	s.logger.Infof("burn from initiated: contract=%s, from=%s, amount=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), amount.String(), txHash.Hex())

	return &pb.BurnFromERC20Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     req.FromAddress,
		Amount:          req.Amount,
	}, nil
}

// DeployERC20 deploys a new ERC20 token contract.
func (s *ERC20Service) DeployERC20(ctx context.Context, req *pb.DeployERC20Request) (*pb.DeployERC20Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate inputs
	if err := validator.ValidateTokenName(req.Name); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	if err := validator.ValidateTokenSymbol(req.Symbol); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	if err := validator.ValidateDecimals(req.Decimals); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Parse initial supply
	initialSupply := big.NewInt(0)
	if req.InitialSupply != "" {
		var err error
		initialSupply, err = validator.ValidateAmount(req.InitialSupply, "initial_supply")
		if err != nil {
			return nil, errors.ToGRPCError(validator.ToAppError(err))
		}
	}

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Get deployer address
	deployerAddr, err := s.contractClient.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get Ethereum client
	client := eth.GetClient()
	if client == nil {
		return nil, errors.ToGRPCError(errors.ErrClientNotInitialized)
	}

	// Determine contract type and owner address
	contractType := getContractType(req.GetContractType())
	ownerAddr := deployerAddr

	// For ownable contracts, use admin address as owner if requested
	if contractType == contract.ContractTypeOwnable && req.GetUseAdmin() {
		adminAddr := keystore.GetAdminAddress()
		if adminAddr == (common.Address{}) {
			return nil, errors.ToGRPCError(errors.InvalidArgument("admin keystore not initialized, cannot use admin address"))
		}
		ownerAddr = adminAddr
		s.logger.Infof("using admin address as owner: %s", ownerAddr.Hex())
	}

	// Deploy contract based on type
	var contractAddr common.Address
	var tx *types.Transaction

	if contractType == contract.ContractTypeOwnable {
		contractAddr, tx, _, err = erc20.DeployERC20TokenOwnable(
			auth,
			client,
			req.Name,
			req.Symbol,
			uint8(req.Decimals),
			initialSupply,
			ownerAddr,
		)
	} else {
		contractAddr, tx, _, err = erc20.DeployERC20Token(
			auth,
			client,
			req.Name,
			req.Symbol,
			uint8(req.Decimals),
			initialSupply,
			ownerAddr,
		)
	}

	if err != nil {
		s.logger.Errorf("failed to deploy ERC20 contract: type=%s, name=%s, symbol=%s, decimals=%d, error=%v",
			contractType, req.Name, req.Symbol, req.Decimals, err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to deploy ERC20 contract"))
	}

	txHash := tx.Hash()
	s.logger.Infof("contract deployed: name=%s, symbol=%s, contract=%s, deployer=%s, tx=%s",
		req.Name, req.Symbol, contractAddr.Hex(), deployerAddr.Hex(), txHash.Hex())

	return &pb.DeployERC20Response{
		TxHash:          txHash.Hex(),
		ContractAddress: contractAddr.Hex(),
		DeployerAddress: deployerAddr.Hex(),
		Name:            req.Name,
		Symbol:          req.Symbol,
		Decimals:        req.Decimals,
		InitialSupply:   req.InitialSupply,
	}, nil
}
