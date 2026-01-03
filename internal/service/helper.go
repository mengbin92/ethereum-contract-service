// Package service provides business logic services for contract interactions.
package service

import (
	"context"

	"eth-contract-service/internal/contract"
	"eth-contract-service/provider/eth"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

// BaseService provides common functionality for all contract services
type BaseService struct {
	logger         *log.Helper
	contractClient *contract.Client
	validator      *ValidationHelper
}

// NewBaseService creates a new base service instance
func NewBaseService(logger log.Logger) *BaseService {
	loggerHelper := log.NewHelper(logger)
	contractClient := contract.NewClient(logger)
	return &BaseService{
		logger:         loggerHelper,
		contractClient: contractClient,
		validator:      NewValidationHelper(),
	}
}

// LogInfo logs an info message with formatted fields
func (s *BaseService) LogInfo(message string, args ...any) {
	s.logger.Infof(message, args...)
}

// LogError logs an error message with formatted fields
func (s *BaseService) LogError(err error, message string, args ...any) {
	if err == nil {
		return
	}
	s.logger.Errorf(message+": %v", append(args, err)...)
}

// LogWarn logs a warning message with formatted fields
func (s *BaseService) LogWarn(message string, args ...any) {
	s.logger.Warnf(message, args...)
}

// WaitForTransaction waits for a transaction to be mined
func (s *BaseService) WaitForTransaction(tx *types.Transaction, contractAddr, operation string) error {
	if tx == nil {
		return errors.New("transaction is nil")
	}

	s.logger.Infof("waiting for transaction to be mined: tx_hash=%s, contract=%s, operation=%s",
		tx.Hash().Hex(), contractAddr, operation)

	// Wait for transaction receipt
	receipt, err := eth.WaitMined(nil, tx.Hash())
	if err != nil {
		s.logger.Errorf("failed to wait for transaction: tx_hash=%s, error=%v",
			tx.Hash().Hex(), err)
		return errors.Wrap(err, "failed to wait for transaction")
	}

	s.logger.Infof("transaction mined successfully: tx_hash=%s, block_number=%s, gas_used=%d, status=%d",
		tx.Hash().Hex(), receipt.BlockNumber.String(), receipt.GasUsed, receipt.Status)

	return nil
}

// CreateTransactOpts creates transaction options with proper error handling
func (s *BaseService) CreateTransactOpts(ctx context.Context, privateKey []byte) (*bind.TransactOpts, error) {
	return s.contractClient.CreateTransactOpts(ctx, privateKey)
}

// ValidateClient validates that the Ethereum client is initialized
func (s *BaseService) ValidateClient() error {
	return s.contractClient.ValidateClient()
}

// GetAddressFromPrivateKey derives the Ethereum address from a private key
func (s *BaseService) GetAddressFromPrivateKey(privateKey []byte) (common.Address, error) {
	return s.contractClient.GetAddressFromPrivateKey(privateKey)
}
