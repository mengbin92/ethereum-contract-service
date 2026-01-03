// Package service provides business logic services for ERC721 token interactions.
package service

import (
	"context"
	"math/big"

	pb "eth-contract-service/api/erc721/v1"
	"eth-contract-service/internal/contract"
	"eth-contract-service/internal/errors"
	"eth-contract-service/internal/validator"
	"eth-contract-service/provider/contract/erc721"
	"eth-contract-service/provider/eth"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kratos/kratos/v2/log"
)

// ERC721Service implements the ERC721 API service.
// It provides methods for interacting with ERC721 (NFT) tokens.
type ERC721Service struct {
	pb.UnimplementedERC721Server
	logger         *log.Helper        // logger for service logging
	contractClient *contract.Client   // client for contract interactions
}

// NewERC721Service creates a new instance of ERC721Service.
func NewERC721Service(logger log.Logger) *ERC721Service {
	return &ERC721Service{
		logger:         log.NewHelper(logger),
		contractClient: contract.NewClient(logger),
	}
}

// GetERC721Balance returns the ERC721 token balance (number of NFTs) of the specified address.
func (s *ERC721Service) GetERC721Balance(ctx context.Context, req *pb.GetERC721BalanceRequest) (*pb.GetERC721BalanceResponse, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get balance
	balance, err := token.BalanceOf(nil, ownerAddr)
	if err != nil {
		s.logger.Errorf("failed to get balance: contract=%s, owner=%s, error=%v", contractAddr.Hex(), ownerAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get balance"))
	}

	s.logger.Infof("balance queried: contract=%s, owner=%s, balance=%s", contractAddr.Hex(), ownerAddr.Hex(), balance.String())

	return &pb.GetERC721BalanceResponse{
		Balance:         balance.String(),
		ContractAddress: req.ContractAddress,
		OwnerAddress:    req.OwnerAddress,
	}, nil
}

// GetERC721TokenInfo returns ERC721 token information.
func (s *ERC721Service) GetERC721TokenInfo(ctx context.Context, req *pb.GetERC721TokenInfoRequest) (*pb.GetERC721TokenInfoResponse, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
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

	s.logger.Infof("token info queried: contract=%s, name=%s, symbol=%s", contractAddr.Hex(), name, symbol)

	return &pb.GetERC721TokenInfoResponse{
		Name:            name,
		Symbol:          symbol,
		ContractAddress: req.ContractAddress,
	}, nil
}

// GetERC721TokenURI returns the URI for a specific token.
func (s *ERC721Service) GetERC721TokenURI(ctx context.Context, req *pb.GetERC721TokenURIRequest) (*pb.GetERC721TokenURIResponse, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get token URI
	tokenURI, err := token.TokenURI(nil, tokenID)
	if err != nil {
		s.logger.Errorf("failed to get token URI: contract=%s, token_id=%s, error=%v", contractAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get token URI"))
	}

	s.logger.Infof("token URI queried: contract=%s, token_id=%s, uri=%s", contractAddr.Hex(), tokenID.String(), tokenURI)

	return &pb.GetERC721TokenURIResponse{
		TokenUri:        tokenURI,
		ContractAddress: req.ContractAddress,
		TokenId:         req.TokenId,
	}, nil
}

// GetERC721OwnerOf returns the owner of a specific token.
func (s *ERC721Service) GetERC721OwnerOf(ctx context.Context, req *pb.GetERC721OwnerOfRequest) (*pb.GetERC721OwnerOfResponse, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get owner
	owner, err := token.OwnerOf(nil, tokenID)
	if err != nil {
		s.logger.Errorf("failed to get owner: contract=%s, token_id=%s, error=%v", contractAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get owner"))
	}

	s.logger.Infof("owner queried: contract=%s, token_id=%s, owner=%s", contractAddr.Hex(), tokenID.String(), owner.Hex())

	return &pb.GetERC721OwnerOfResponse{
		OwnerAddress:    owner.Hex(),
		ContractAddress: req.ContractAddress,
		TokenId:         req.TokenId,
	}, nil
}

// GetERC721Approved returns the approved address for a token.
func (s *ERC721Service) GetERC721Approved(ctx context.Context, req *pb.GetERC721ApprovedRequest) (*pb.GetERC721ApprovedResponse, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Get approved address
	approved, err := token.GetApproved(nil, tokenID)
	if err != nil {
		s.logger.Errorf("failed to get approved: contract=%s, token_id=%s, error=%v", contractAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to get approved"))
	}

	s.logger.Infof("approved queried: contract=%s, token_id=%s, approved=%s", contractAddr.Hex(), tokenID.String(), approved.Hex())

	return &pb.GetERC721ApprovedResponse{
		ApprovedAddress: approved.Hex(),
		ContractAddress: req.ContractAddress,
		TokenId:         req.TokenId,
	}, nil
}

// IsApprovedForAllERC721 checks if an operator is approved for all tokens of an owner.
func (s *ERC721Service) IsApprovedForAllERC721(ctx context.Context, req *pb.IsApprovedForAllERC721Request) (*pb.IsApprovedForAllERC721Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate address
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	ownerAddr, err := validator.ValidateAddress(req.OwnerAddress, "owner_address")
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Check approval
	approved, err := token.IsApprovedForAll(nil, ownerAddr, operatorAddr)
	if err != nil {
		s.logger.Errorf("failed to check approval: contract=%s, owner=%s, operator=%s, error=%v", contractAddr.Hex(), ownerAddr.Hex(), operatorAddr.Hex(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to check approval"))
	}

	s.logger.Infof("approval checked: contract=%s, owner=%s, operator=%s, approved=%v", contractAddr.Hex(), ownerAddr.Hex(), operatorAddr.Hex(), approved)

	return &pb.IsApprovedForAllERC721Response{
		Approved:        approved,
		ContractAddress: req.ContractAddress,
		OwnerAddress:    req.OwnerAddress,
		OperatorAddress: req.OperatorAddress,
	}, nil
}

// TransferERC721 transfers an ERC721 token from the caller to the specified address.
func (s *ERC721Service) TransferERC721(ctx context.Context, req *pb.TransferERC721Request) (*pb.TransferERC721Response, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Transfer token
	tx, err := token.TransferFrom(auth, fromAddr, toAddr, tokenID)
	if err != nil {
		s.logger.Errorf("failed to transfer token: contract=%s, from=%s, to=%s, token_id=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to transfer token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("transfer initiated: contract=%s, from=%s, to=%s, token_id=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), txHash.Hex())

	return &pb.TransferERC721Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     fromAddr.Hex(),
		ToAddress:       req.ToAddress,
		TokenId:         req.TokenId,
	}, nil
}

// SafeTransferERC721 safely transfers an ERC721 token.
func (s *ERC721Service) SafeTransferERC721(ctx context.Context, req *pb.SafeTransferERC721Request) (*pb.SafeTransferERC721Response, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Safe transfer token
	tx, err := token.SafeTransferFrom(auth, fromAddr, toAddr, tokenID)
	if err != nil {
		s.logger.Errorf("failed to safe transfer token: contract=%s, from=%s, to=%s, token_id=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to safe transfer token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("safe transfer initiated: contract=%s, from=%s, to=%s, token_id=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), txHash.Hex())

	return &pb.SafeTransferERC721Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     fromAddr.Hex(),
		ToAddress:       req.ToAddress,
		TokenId:         req.TokenId,
	}, nil
}

// SafeTransferERC721WithData safely transfers an ERC721 token with additional data.
func (s *ERC721Service) SafeTransferERC721WithData(ctx context.Context, req *pb.SafeTransferERC721WithDataRequest) (*pb.SafeTransferERC721WithDataResponse, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Safe transfer token with data
	tx, err := token.SafeTransferFrom0(auth, fromAddr, toAddr, tokenID, req.Data)
	if err != nil {
		s.logger.Errorf("failed to safe transfer token with data: contract=%s, from=%s, to=%s, token_id=%s, error=%v",
			contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to safe transfer token with data"))
	}

	txHash := tx.Hash()
	s.logger.Infof("safe transfer with data initiated: contract=%s, from=%s, to=%s, token_id=%s, tx=%s",
		contractAddr.Hex(), fromAddr.Hex(), toAddr.Hex(), tokenID.String(), txHash.Hex())

	return &pb.SafeTransferERC721WithDataResponse{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		FromAddress:     fromAddr.Hex(),
		ToAddress:       req.ToAddress,
		TokenId:         req.TokenId,
	}, nil
}

// ApproveERC721 approves another address to transfer the specified token.
func (s *ERC721Service) ApproveERC721(ctx context.Context, req *pb.ApproveERC721Request) (*pb.ApproveERC721Response, error) {
	if err := validator.ValidateRequest(req); err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate addresses
	contractAddr, err := validator.ValidateContractAddress(req.ContractAddress)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	approvedAddr, err := validator.ValidateAddress(req.ApprovedAddress, "approved_address")
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate token ID
	tokenID, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		return nil, errors.ToGRPCError(errors.InvalidArgument("invalid token_id format"))
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Approve token
	tx, err := token.Approve(auth, approvedAddr, tokenID)
	if err != nil {
		s.logger.Errorf("failed to approve token: contract=%s, owner=%s, approved=%s, token_id=%s, error=%v",
			contractAddr.Hex(), ownerAddr.Hex(), approvedAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to approve token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("approval initiated: contract=%s, owner=%s, approved=%s, token_id=%s, tx=%s",
		contractAddr.Hex(), ownerAddr.Hex(), approvedAddr.Hex(), tokenID.String(), txHash.Hex())

	return &pb.ApproveERC721Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		OwnerAddress:    ownerAddr.Hex(),
		ApprovedAddress: req.ApprovedAddress,
		TokenId:         req.TokenId,
	}, nil
}

// SetApprovalForAllERC721 enables or disables approval for a third party to manage all tokens.
func (s *ERC721Service) SetApprovalForAllERC721(ctx context.Context, req *pb.SetApprovalForAllERC721Request) (*pb.SetApprovalForAllERC721Response, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
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

	return &pb.SetApprovalForAllERC721Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		OwnerAddress:    ownerAddr.Hex(),
		OperatorAddress: req.OperatorAddress,
		Approved:        req.Approved,
	}, nil
}

// SafeMintERC721 safely mints a new ERC721 token.
func (s *ERC721Service) SafeMintERC721(ctx context.Context, req *pb.SafeMintERC721Request) (*pb.SafeMintERC721Response, error) {
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

	// Validate private key
	privateKey, err := validator.ValidatePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, errors.ToGRPCError(validator.ToAppError(err))
	}

	// Validate client
	if err := s.contractClient.ValidateClient(); err != nil {
		return nil, errors.ToGRPCError(err)
	}

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Mint token
	tx, err := token.SafeMint(auth, toAddr, tokenID)
	if err != nil {
		s.logger.Errorf("failed to mint token: contract=%s, to=%s, token_id=%s, error=%v",
			contractAddr.Hex(), toAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to mint token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("mint initiated: contract=%s, to=%s, token_id=%s, tx=%s",
		contractAddr.Hex(), toAddr.Hex(), tokenID.String(), txHash.Hex())

	return &pb.SafeMintERC721Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		ToAddress:       req.ToAddress,
		TokenId:         req.TokenId,
	}, nil
}

// BurnERC721 burns an ERC721 token.
func (s *ERC721Service) BurnERC721(ctx context.Context, req *pb.BurnERC721Request) (*pb.BurnERC721Response, error) {
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

	// Create ERC721 contract instance
	token, err := s.contractClient.GetERC721Token(contractAddr)
	if err != nil {
		s.logger.Errorf("failed to get ERC721 token: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Create transaction options
	auth, err := s.contractClient.CreateTransactOpts(ctx, privateKey)
	if err != nil {
		s.logger.Errorf("failed to create transaction options: %v", err)
		return nil, errors.ToGRPCError(err)
	}

	// Burn token
	tx, err := token.Burn(auth, tokenID)
	if err != nil {
		s.logger.Errorf("failed to burn token: contract=%s, token_id=%s, error=%v",
			contractAddr.Hex(), tokenID.String(), err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to burn token"))
	}

	txHash := tx.Hash()
	s.logger.Infof("burn initiated: contract=%s, owner=%s, token_id=%s, tx=%s",
		contractAddr.Hex(), ownerAddr.Hex(), tokenID.String(), txHash.Hex())

	return &pb.BurnERC721Response{
		TxHash:          txHash.Hex(),
		ContractAddress: req.ContractAddress,
		TokenId:         req.TokenId,
	}, nil
}

// DeployERC721 deploys a new ERC721 token contract.
func (s *ERC721Service) DeployERC721(ctx context.Context, req *pb.DeployERC721Request) (*pb.DeployERC721Response, error) {
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
	contractAddr, tx, _, err := s.deployERC721Token(auth, ownerAddr, req.Name, req.Symbol)
	if err != nil {
		s.logger.Errorf("failed to deploy ERC721 contract: name=%s, symbol=%s, error=%v",
			req.Name, req.Symbol, err)
		return nil, errors.ToGRPCError(errors.WrapError(err, errors.CodeInternal, "failed to deploy ERC721 contract"))
	}

	txHash := tx.Hash()
	s.logger.Infof("contract deployed: name=%s, symbol=%s, contract=%s, deployer=%s, tx=%s",
		req.Name, req.Symbol, contractAddr.Hex(), deployerAddr.Hex(), txHash.Hex())

	return &pb.DeployERC721Response{
		TxHash:          txHash.Hex(),
		ContractAddress: contractAddr.Hex(),
		DeployerAddress: deployerAddr.Hex(),
		Name:            req.Name,
		Symbol:          req.Symbol,
	}, nil
}

// deployERC721Token is a helper function to deploy ERC721 token
func (s *ERC721Service) deployERC721Token(auth *bind.TransactOpts, initialOwner common.Address, name string, symbol string) (common.Address, *types.Transaction, *erc721.Erc721, error) {
	client := eth.GetClient()
	if client == nil {
		return common.Address{}, nil, nil, errors.ErrClientNotInitialized
	}

	return erc721.DeployErc721(auth, client, initialOwner, name, symbol)
}
