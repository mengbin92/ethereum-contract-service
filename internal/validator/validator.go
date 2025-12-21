// Package validator provides validation utilities for Ethereum-related inputs.
package validator

import (
	"encoding/hex"
	"math/big"
	"strings"

	"eth-contract-service/internal/errors"

	"github.com/ethereum/go-ethereum/common"
	pkgErrors "github.com/pkg/errors"
)

// ValidateAddress validates an Ethereum address
func ValidateAddress(addr string, fieldName string) (common.Address, error) {
	if addr == "" {
		return common.Address{}, pkgErrors.Errorf("%s cannot be empty", fieldName)
	}

	if !common.IsHexAddress(addr) {
		return common.Address{}, pkgErrors.Errorf("%s is not a valid Ethereum address: %s", fieldName, addr)
	}

	return common.HexToAddress(addr), nil
}

// ValidateContractAddress validates a contract address
func ValidateContractAddress(addr string) (common.Address, error) {
	return ValidateAddress(addr, "contract_address")
}

// ValidateAmount validates and parses an amount string into a big.Int
func ValidateAmount(amountStr string, fieldName string) (*big.Int, error) {
	if amountStr == "" {
		return nil, pkgErrors.Errorf("%s cannot be empty", fieldName)
	}

	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return nil, pkgErrors.Errorf("invalid %s: %s (must be a decimal number)", fieldName, amountStr)
	}

	if amount.Sign() < 0 {
		return nil, pkgErrors.Errorf("%s cannot be negative", fieldName)
	}

	return amount, nil
}

// ValidatePrivateKey validates and parses a hex-encoded private key
//
// Supported formats:
//   - With 0x prefix: "0x1234567890abcdef..." (64 hex characters after 0x)
//   - Without 0x prefix: "1234567890abcdef..." (64 hex characters)
//   - Case insensitive: "0X..." or "0x..." are both accepted
//
// Returns:
//   - []byte: 32-byte private key
//   - error: Error if validation fails
func ValidatePrivateKey(keyStr string) ([]byte, error) {
	if keyStr == "" {
		return nil, pkgErrors.New("private_key cannot be empty")
	}

	// Remove 0x or 0X prefix if present
	keyStr = strings.TrimPrefix(keyStr, "0x")
	keyStr = strings.TrimPrefix(keyStr, "0X")

	// Decode hex string to bytes
	privateKey, err := hex.DecodeString(keyStr)
	if err != nil {
		return nil, pkgErrors.Wrapf(err, "failed to decode private key: invalid hex format (expected 64 hex characters)")
	}

	// Validate length: Ethereum private keys must be exactly 32 bytes
	if len(privateKey) != 32 {
		return nil, pkgErrors.Errorf("private key must be 32 bytes (64 hex characters), got %d bytes (%d hex characters)", len(privateKey), len(keyStr))
	}

	return privateKey, nil
}

// ValidateDecimals validates token decimals
func ValidateDecimals(decimals uint32) error {
	if decimals == 0 || decimals > 18 {
		return pkgErrors.New("decimals must be between 1 and 18")
	}
	return nil
}

// ValidateTokenName validates token name
func ValidateTokenName(name string) error {
	if name == "" {
		return pkgErrors.New("name cannot be empty")
	}
	if len(name) > 100 {
		return pkgErrors.New("name cannot exceed 100 characters")
	}
	return nil
}

// ValidateTokenSymbol validates token symbol
func ValidateTokenSymbol(symbol string) error {
	if symbol == "" {
		return pkgErrors.New("symbol cannot be empty")
	}
	if len(symbol) > 20 {
		return pkgErrors.New("symbol cannot exceed 20 characters")
	}
	return nil
}

// ValidateRequest validates a request is not nil
func ValidateRequest(req interface{}) error {
	if req == nil {
		return pkgErrors.New("request cannot be nil")
	}
	return nil
}

// ToAppError converts a validation error to an AppError
func ToAppError(err error) *errors.AppError {
	if err == nil {
		return nil
	}

	if appErr, ok := err.(*errors.AppError); ok {
		return appErr
	}

	return errors.InvalidArgument("%s", err.Error())
}
