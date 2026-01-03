// Package service provides business logic services for ERC1155 token interactions.
package service

import (
	"context"
	"math/big"

	pb "eth-contract-service/api/erc1155/v1"
	"eth-contract-service/internal/contract"
	"eth-contract-service/internal/errors"
	"eth-contract-service/internal/validator"
	"eth-contract-service/provider/contract/erc1155"
	"eth-contract-service/provider/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kratos/kratos/v2/log"
)

// ERC1155Service implements the ERC1155 API service.
// It provides methods for interacting with ERC1155 (Multi-Token) tokens.
type ERC1155Service struct {
	pb.UnimplementedERC1155Server
	logger         *log.Helper        // logger for service logging
	contractClient *contract.Client   // client for contract interactions
}

// NewERC1155Service creates a new instance of ERC1155Service.
func NewERC1155Service(logger log.Logger) *ERC1155Service {
	return &ERC1155Service{
		logger:         log.NewHelper(logger),
		contractClient: contract.NewClient(logger),
	}
}

// GetERC1155Balance returns the balance of an address for a specific token ID.
func (s *ERC1155Service) GetERC1155Balance(ctx context.Context, req *pb.GetERC1155BalanceRequest) (*pb.GetERC1155BalanceResponse, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	accountAddr, err := validator.ValidateAddress(req.AccountAddress, "account_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate token ID
	tokenID, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format"))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get balance
	balance, err := token.BalanceOf(nil, accountAddr, tokenID)
	if err != nil {
		s.logger.Errorf("failed to get balance: contract=%s, account=%s, token_id=%s, error=%v",
			contractAddr.Hex(), accountAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get balance"))
	}

	s.logger.Infof("balance queried: contract=%s, account=%s, token_id=%s, balance=%s",
		contractAddr.Hex(), accountAddr.Hex(), tokenID.String(), balance.String())

	return &pb.GetERC1155BalanceResponse{
		Balance:         balance.String(),
		ContractAddress: req.ContractAddress,
		AccountAddress:  req.AccountAddress,
		TokenId:         req.TokenId,
	}, nil
}

// GetERC1155BalancesBatch returns the balance of multiple addresses for multiple token IDs.
func (s *ERC1155Service) GetERC1155BalancesBatch(ctx context.Context, req *pb.GetERC1155BalancesBatchRequest) (*pb.GetERC1155BalancesBatchResponse, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate address
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate accounts array
	if len(req.Accounts) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("accounts array cannot be empty"))
	}

	// Validate token IDs array
	if len(req.TokenIds) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("token_ids array cannot be empty"))
	}

	if len(req.Accounts) != len(req.TokenIds) {
		return nil, errors.ToGRPCError(errors.InvalidArgument("accounts and token_ids arrays must have the same length"))
	}

	// Convert string addresses to common.Address
	accounts := make([]common.Address, len(req.Accounts))
	for i, addrStr := range req.Accounts {
		addr, err := validator.ValidateAddress(addrStr, "account")
		if err != nil {
			return nil, errors.ToGRPCError(validator.ToAppError(err))
		}
		accounts[i] = addr
	}

	// Convert string token IDs to *big.Int
	tokenIDs := make([]*big.Int, len(req.TokenIds))
	for i, idStr := range req.TokenIds {
		tokenID, ok := new(big.Int).SetString(idStr, 10)
		if !ok {
			return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format at index %d", i))
		}
		tokenIDs[i] = tokenID
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get batch balances
	balances, err := token.BalanceOfBatch(nil, accounts, tokenIDs)
	if err != nil {
		s.logger.Errorf("failed to get batch balances: contract=%s, error=%v", contractAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get batch balances"))
	}

	// Convert balances to string array
	balanceStrings := make([]string, len(balances))
	for i, balance := range balances {
		balanceStrings[i] = balance.String()
	}

	s.logger.Infof("batch balances queried: contract=%s, num_accounts=%d", contractAddr.Hex(), len(accounts))

	return &pb.GetERC1155BalancesBatchResponse{
		Balances:        balanceStrings,
		ContractAddress: req.ContractAddress,
		Accounts:        req.Accounts,
		TokenIds:        req.TokenIds,
	}, nil
}

// GetERC1155TokenURI returns the URI for a specific token ID.
func (s *ERC1155Service) GetERC1155TokenURI(ctx context.Context, req *pb.GetERC1155TokenURIRequest) (*pb.GetERC1155TokenURIResponse, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate address
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate token ID
	tokenID, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format"))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get token URI
	tokenURI, err := token.Uri(nil, tokenID)
	if err != nil {
		s.logger.Errorf("failed to get token URI: contract=%s, token_id=%s, error=%v", contractAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get token URI"))
	}

	s.logger.Infof("token URI queried: contract=%s, token_id=%s, uri=%s", contractAddr.Hex(), tokenID.String(), tokenURI)

	return &pb.GetERC1155TokenURIResponse{
		TokenUri:        tokenURI,
		ContractAddress: req.ContractAddress,
		TokenId:         req.TokenId,
	}, nil
}

// IsApprovedForAllERC1155 checks if an operator is approved for all tokens of an owner.
func (s *ERC1155Service) IsApprovedForAllERC1155(ctx context.Context, req *pb.IsApprovedForAllERC1155Request) (*pb.IsApprovedForAllERC1155Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate address
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	accountAddr, err := validator.ValidateAddress(req.AccountAddress, "account_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	operatorAddr, err := validator.ValidateAddress(req.OperatorAddress, "operator_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Check approval
	approved, err := token.IsApprovedForAll(nil, accountAddr, operatorAddr)
	if err != nil {
		s.logger.Errorf("failed to check approval: contract=%s, account=%s, operator=%s, error=%v",
			contractAddr.Hex(), accountAddr.Hex(), operatorAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to check approval"))
	}

	s.logger.Infof("approval checked: contract=%s, account=%s, operator=%s, approved=%v",
		contractAddr.Hex(), accountAddr.Hex(), operatorAddr.Hex(), approved)

	return &pb.IsApprovedForAllERC1155Response{
		Approved:        approved,
		ContractAddress: req.ContractAddress,
		AccountAddress:  req.AccountAddress,
		OperatorAddress: req.OperatorAddress,
	}, nil
}

// SafeTransferERC1155 transfers an ERC1155 token from one address to another.
func (s *ERC1155Service) SafeTransferERC1155(ctx context.Context, req *pb.SafeTransferERC1155Request) (*pb.SafeTransferERC1155Response, error) {
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

	// Validate token ID
	tokenID, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format"))
	}

	// Validate amount
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid amount format"))
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
	senderAddr, err := s.contractClient.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Verify sender matches from_address
	if senderAddr.Hex() != fromAddr.Hex() {
		return nil, errors.ToGRPCError(errors.InvalidArgument("private key does not match from_address"))
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Transfer token
	tx, err := token.SafeTransferFrom(auth, fromAddr, toAddr, tokenID, amount, req.Data)
	if err != nil {
		s.logger.Errorf("failed to transfer token: contract=%s, from=%s, to=%s, token_id=%s, amount=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to transfer token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("transfer initiated: contract=%s, from=%s, to=%s, token_id=%s, amount=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), amount.String(), txHash.Hex())

	return &pb.SafeTransferERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     fromAddr.Hex(),
		ToAddress:       req.ToAddress,
		TokenId:         req.TokenId,
		Amount:          req.Amount,
	}, nil
}

// SafeBatchTransferERC1155 safely transfers multiple ERC1155 tokens.
func (s *ERC1155Service) SafeBatchTransferERC1155(ctx context.Context, req *pb.SafeBatchTransferERC1155Request) (*pb.SafeBatchTransferERC1155Response, error) {
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

	// Validate token IDs array
	if len(req.TokenIds) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("token_ids array cannot be empty"))
	}

	// Validate amounts array
	if len(req.Amounts) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("amounts array cannot be empty"))
	}

	if len(req.TokenIds) != len(req.Amounts) {
		return nil, errors.ToGRPCError(errors.InvalidArgument("token_ids and amounts arrays must have the same length"))
	}

	// Convert string token IDs to *big.Int
	tokenIDs := make([]*big.Int, len(req.TokenIds))
	for i, idStr := range req.TokenIds {
		tokenID, ok := new(big.Int).SetString(idStr, 10)
		if !ok {
			return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format at index %d", i))
		}
		tokenIDs[i] = tokenID
	}

	// Convert string amounts to *big.Int
	amounts := make([]*big.Int, len(req.Amounts))
	for i, amountStr := range req.Amounts {
		amount, ok := new(big.Int).SetString(amountStr, 10)
		if !ok {
			return nil, errors.ToGRPCError(errors.InvalidArgument("invalid amount format at index %d", i))
		}
		amounts[i] = amount
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
	senderAddr, err := s.contractClient.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Verify sender matches from_address
	if senderAddr.Hex() != fromAddr.Hex() {
		return nil, errors.ToGRPCError(errors.InvalidArgument("private key does not match from_address"))
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Batch transfer tokens
	tx, err := token.SafeBatchTransferFrom(auth, fromAddr, toAddr, tokenIDs, amounts, req.Data)
	if err != nil {
		s.logger.Errorf("failed to batch transfer tokens: contract=%s, from=%s, to=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to batch transfer tokens"))
	}

	txHash := tx.Hash()
	s.logger.Infof("batch transfer initiated: contract=%s, from=%s, to=%s, num_tokens=%d, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), len(tokenIDs), txHash.Hex())

	return &pb.SafeBatchTransferERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     fromAddr.Hex(),
		ToAddress:       req.ToAddress,
		TokenIds:        req.TokenIds,
		Amounts:         req.Amounts,
	}, nil
}

// SetApprovalForAllERC1155 enables or disables approval for a third party to manage all tokens.
func (s *ERC1155Service) SetApprovalForAllERC1155(ctx context.Context, req *pb.SetApprovalForAllERC1155Request) (*pb.SetApprovalForAllERC1155Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	operatorAddr, err := validator.ValidateAddress(req.OperatorAddress, "operator_address")
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

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Set approval for all
	tx, err := token.SetApprovalForAll(auth, operatorAddr, req.Approved)
	if err != nil {
		s.logger.Errorf("failed to set approval for all: contract=%s, owner=%s, operator=%s, approved=%v, error=%v",
			contractAddr.Hex(), ownerAddr.Hex(), operatorAddr.Hex(), req.Approved, err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to set approval for all"))
	}

	txHash := tx.Hash()
	s.logger.Infof("approval for all initiated: contract=%s, owner=%s, operator=%s, approved=%v, tx=%s",
		contractAddr.Hex(), ownerAddr.Hex(), operatorAddr.Hex(), req.Approved, txHash.Hex())

	return &pb.SetApprovalForAllERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		OwnerAddress:    ownerAddr.Hex(),
		OperatorAddress: req.OperatorAddress,
		Approved:        req.Approved,
	}, nil
}

// MintERC1155 mints new ERC1155 tokens.
func (s *ERC1155Service) MintERC1155(ctx context.Context, req *pb.MintERC1155Request) (*pb.MintERC1155Response, error) {
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

	// Validate token ID
	tokenID, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format"))
	}

	// Validate amount
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid amount format"))
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

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Mint token
	tx, err := token.Mint(auth, toAddr, tokenID, amount, req.Data)
	if err != nil {
		s.logger.Errorf("failed to mint token: contract=%s, to=%s, token_id=%s, amount=%s, error=%v",
			contractAddr.Hex(), toAddr.Hex(), tokenID.String(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to mint token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("mint initiated: contract=%s, to=%s, token_id=%s, amount=%s, tx=%s",
		contractAddr.Hex(), toAddr.Hex(), tokenID.String(), amount.String(), txHash.Hex())

	return &pb.MintERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		ToAddress:       req.ToAddress,
		TokenId:         req.TokenId,
		Amount:          req.Amount,
	}, nil
}

// MintBatchERC1155 mints multiple ERC1155 tokens.
func (s *ERC1155Service) MintBatchERC1155(ctx context.Context, req *pb.MintBatchERC1155Request) (*pb.MintBatchERC1155Response, error) {
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

	// Validate token IDs array
	if len(req.TokenIds) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("token_ids array cannot be empty"))
	}

	// Validate amounts array
	if len(req.Amounts) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("amounts array cannot be empty"))
	}

	if len(req.TokenIds) != len(req.Amounts) {
		return nil, errors.ToGRPCError(errors.InvalidArgument("token_ids and amounts arrays must have the same length"))
	}

	// Convert string token IDs to *big.Int
	tokenIDs := make([]*big.Int, len(req.TokenIds))
	for i, idStr := range req.TokenIds {
		tokenID, ok := new(big.Int).SetString(idStr, 10)
		if !ok {
			return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format at index %d", i))
		}
		tokenIDs[i] = tokenID
	}

	// Convert string amounts to *big.Int
	amounts := make([]*big.Int, len(req.Amounts))
	for i, amountStr := range req.Amounts {
		amount, ok := new(big.Int).SetString(amountStr, 10)
		if !ok {
			return nil, errors.ToGRPCError(errors.InvalidArgument("invalid amount format at index %d", i))
		}
		amounts[i] = amount
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

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Batch mint tokens
	tx, err := token.MintBatch(auth, toAddr, tokenIDs, amounts, req.Data)
	if err != nil {
		s.logger.Errorf("failed to batch mint tokens: contract=%s, to=%s, error=%v",
			contractAddr.Hex(), toAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to batch mint tokens"))
	}

	txHash := tx.Hash()
	s.logger.Infof("batch mint initiated: contract=%s, to=%s, num_tokens=%d, tx=%s",
		contractAddr.Hex(), toAddr.Hex(), len(tokenIDs), txHash.Hex())

	return &pb.MintBatchERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		ToAddress:       req.ToAddress,
		TokenIds:        req.TokenIds,
		Amounts:         req.Amounts,
	}, nil
}

// BurnERC1155 burns ERC1155 tokens.
func (s *ERC1155Service) BurnERC1155(ctx context.Context, req *pb.BurnERC1155Request) (*pb.BurnERC1155Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	accountAddr, err := validator.ValidateAddress(req.AccountAddress, "account_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate token ID
	tokenID, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format"))
	}

	// Validate amount
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid amount format"))
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

	// Verify owner matches account_address
	if ownerAddr.Hex() != accountAddr.Hex() {
		return nil, errors.ToGRPCError(errors.InvalidArgument("private key does not match account_address"))
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Burn token
	tx, err := token.Burn(auth, accountAddr, tokenID, amount)
	if err != nil {
		s.logger.Errorf("failed to burn token: contract=%s, account=%s, token_id=%s, amount=%s, error=%v",
			contractAddr.Hex(), accountAddr.Hex(), tokenID.String(), amount.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to burn token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("burn initiated: contract=%s, account=%s, token_id=%s, amount=%s, tx=%s",
		contractAddr.Hex(), accountAddr.Hex(), tokenID.String(), amount.String(), txHash.Hex())

	return &pb.BurnERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		AccountAddress:  req.AccountAddress,
		TokenId:         req.TokenId,
		Amount:          req.Amount,
	}, nil
}

// BurnBatchERC1155 burns multiple ERC1155 tokens.
func (s *ERC1155Service) BurnBatchERC1155(ctx context.Context, req *pb.BurnBatchERC1155Request) (*pb.BurnBatchERC1155Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	accountAddr, err := validator.ValidateAddress(req.AccountAddress, "account_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate token IDs array
	if len(req.TokenIds) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("token_ids array cannot be empty"))
	}

	// Validate amounts array
	if len(req.Amounts) == 0 {
		return nil, errors.ToGRPCError(errors.InvalidArgument("amounts array cannot be empty"))
	}

	if len(req.TokenIds) != len(req.Amounts) {
		return nil, errors.ToGRPCError(errors.InvalidArgument("token_ids and amounts arrays must have the same length"))
	}

	// Convert string token IDs to *big.Int
	tokenIDs := make([]*big.Int, len(req.TokenIds))
	for i, idStr := range req.TokenIds {
		tokenID, ok := new(big.Int).SetString(idStr, 10)
		if !ok {
			return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format at index %d", i))
		}
		tokenIDs[i] = tokenID
	}

	// Convert string amounts to *big.Int
	amounts := make([]*big.Int, len(req.Amounts))
	for i, amountStr := range req.Amounts {
		amount, ok := new(big.Int).SetString(amountStr, 10)
		if !ok {
			return nil, errors.ToGRPCError(errors.InvalidArgument("invalid amount format at index %d", i))
		}
		amounts[i] = amount
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

	// Verify owner matches account_address
	if ownerAddr.Hex() != accountAddr.Hex() {
		return nil, errors.ToGRPCError(errors.InvalidArgument("private key does not match account_address"))
	}

	// Create ERC1155 contract instance
	token, err := s.contractClient.GetERC1155Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC1155 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Batch burn tokens
	tx, err := token.BurnBatch(auth, accountAddr, tokenIDs, amounts)
	if err != nil {
		s.logger.Errorf("failed to batch burn tokens: contract=%s, account=%s, error=%v",
			contractAddr.Hex(), accountAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to batch burn tokens"))
	}

	txHash := tx.Hash()
	s.logger.Infof("batch burn initiated: contract=%s, account=%s, num_tokens=%d, tx=%s",
		contractAddr.Hex(), accountAddr.Hex(), len(tokenIDs), txHash.Hex())

	return &pb.BurnBatchERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		AccountAddress:  req.AccountAddress,
		TokenIds:        req.TokenIds,
		Amounts:         req.Amounts,
	}, nil
}

// DeployERC1155 deploys a new ERC1155 token contract.
func (s *ERC1155Service) DeployERC1155(ctx context.Context, req *pb.DeployERC1155Request) (*pb.DeployERC1155Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate URI
	if req.Uri == "" {
		return nil, errors.ToGRPCError(errors.InvalidArgument("uri cannot be empty"))
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

	// Determine initial owner
	ownerAddr := deployerAddr
	if req.InitialOwner != "" {
		ownerAddr, err = validator.ValidateAddress(req.InitialOwner, "initial_owner")
		if err != nil {
			return nil, errors.ToGRPCError(validator.ToAppError(err))
		}
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Deploy contract
	contractAddr, tx, _, err := erc1155.DeployErc1155(auth, eth.GetClient(), ownerAddr, req.Uri)
	if err != nil {
		s.logger.Errorf("failed to deploy ERC1155 contract: uri=%s, error=%v", req.Uri, err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to deploy ERC1155 contract"))
	}

	txHash := tx.Hash()
	s.logger.Infof("contract deployed: uri=%s, contract=%s, deployer=%s, tx=%s",
		req.Uri, contractAddr.Hex(), deployerAddr.Hex(), txHash.Hex())

	return &pb.DeployERC1155Response{
		TxHash:          txHash.Hex(),
		ContractAddress: contractAddr.Hex(),
		DeployerAddress: deployerAddr.Hex(),
		Uri:             req.Uri,
	}, nil
}
