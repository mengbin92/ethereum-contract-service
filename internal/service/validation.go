// Package service provides business logic services for contract interactions.
package service

import (
	"math/big"

	"eth-contract-service/internal/validator"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// ValidationHelper provides helper functions for common validation operations
type ValidationHelper struct{}

// NewValidationHelper creates a new validation helper
func NewValidationHelper() *ValidationHelper {
	return &ValidationHelper{}
}

// ValidateContractAddress validates a contract address
func (v *ValidationHelper) ValidateContractAddress(contractAddress string, fieldName string) (common.Address, error) {
	addr, err := validator.ValidateContractAddress(contractAddress)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "invalid contract address")
	}
	return addr, nil
}

// ValidateAddress validates an Ethereum address
func (v *ValidationHelper) ValidateAddress(address string, fieldName string) (common.Address, error) {
	addr, err := validator.ValidateAddress(address, fieldName)
	if err != nil {
		return common.Address{}, errors.Wrapf(err, "invalid %s", fieldName)
	}
	return addr, nil
}

// ValidateAmount validates and parses an amount string
func (v *ValidationHelper) ValidateAmount(amount string, fieldName string) (*big.Int, error) {
	amt, err := validator.ValidateAmount(amount, fieldName)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid %s", fieldName)
	}
	return amt, nil
}

// ValidatePrivateKey validates a private key
func (v *ValidationHelper) ValidatePrivateKey(privateKey string) ([]byte, error) {
	key, err := validator.ValidatePrivateKey(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "invalid private key")
	}
	return key, nil
}

// ValidateTokenID validates a token ID string
func (v *ValidationHelper) ValidateTokenID(tokenID string, fieldName string) (*big.Int, error) {
	id, ok := new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, errors.Errorf("invalid %s format", fieldName)
	}
	return id, nil
}

// ValidateTokenName validates a token name
func (v *ValidationHelper) ValidateTokenName(name string) error {
	return validator.ValidateTokenName(name)
}

// ValidateTokenSymbol validates a token symbol
func (v *ValidationHelper) ValidateTokenSymbol(symbol string) error {
	return validator.ValidateTokenSymbol(symbol)
}

// ValidateDecimals validates decimals
func (v *ValidationHelper) ValidateDecimals(decimals uint32) error {
	return validator.ValidateDecimals(decimals)
}

// ValidateAddresses validates multiple addresses
func (v *ValidationHelper) ValidateAddresses(addresses []string, fieldNamePrefix string) ([]common.Address, error) {
	validatedAddrs := make([]common.Address, len(addresses))
	for i, addrStr := range addresses {
		addr, err := v.ValidateAddress(addrStr, fieldNamePrefix)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid %s at index %d", fieldNamePrefix, i)
		}
		validatedAddrs[i] = addr
	}
	return validatedAddrs, nil
}

// ValidateTokenIDs validates multiple token IDs
func (v *ValidationHelper) ValidateTokenIDs(tokenIDs []string, fieldNamePrefix string) ([]*big.Int, error) {
	validatedIDs := make([]*big.Int, len(tokenIDs))
	for i, idStr := range tokenIDs {
		id, err := v.ValidateTokenID(idStr, fieldNamePrefix)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid %s at index %d", fieldNamePrefix, i)
		}
		validatedIDs[i] = id
	}
	return validatedIDs, nil
}

// ValidateAmounts validates multiple amounts
func (v *ValidationHelper) ValidateAmounts(amounts []string, fieldNamePrefix string) ([]*big.Int, error) {
	validatedAmounts := make([]*big.Int, len(amounts))
	for i, amountStr := range amounts {
		amount, err := v.ValidateAmount(amountStr, fieldNamePrefix)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid %s at index %d", fieldNamePrefix, i)
		}
		validatedAmounts[i] = amount
	}
	return validatedAmounts, nil
}

// ValidateArraysSameLength validates that two arrays have the same length
func (v *ValidationHelper) ValidateArraysSameLength(arr1, arr2 []string, name1, name2 string) error {
	if len(arr1) != len(arr2) {
		return errors.Errorf("%s and %s arrays must have the same length", name1, name2)
	}
	return nil
}

// ValidateArrayNotEmpty validates that an array is not empty
func (v *ValidationHelper) ValidateArrayNotEmpty(arr []string, fieldName string) error {
	if len(arr) == 0 {
		return errors.Errorf("%s cannot be empty", fieldName)
	}
	return nil
}
