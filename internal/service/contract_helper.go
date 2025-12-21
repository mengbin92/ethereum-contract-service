// Package service provides business logic services for ERC20 token interactions.
package service

import (
	"math/big"

	"eth-contract-service/internal/contract"
	"eth-contract-service/provider/contract/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC20Contract interface defines common methods for both ERC20Token and ERC20TokenOwnable
type ERC20Contract interface {
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
	Name(opts *bind.CallOpts) (string, error)
	Symbol(opts *bind.CallOpts) (string, error)
	Decimals(opts *bind.CallOpts) (uint8, error)
	TotalSupply(opts *bind.CallOpts) (*big.Int, error)
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)
	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)
	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)
	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)
	Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)
	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)
	BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error)
}

// getContractType parses contract type from string
//
//nolint:unused // This function is used in getERC20Contract
func getContractType(contractTypeStr string) contract.ContractType {
	if contractTypeStr == "ownable" {
		return contract.ContractTypeOwnable
	}
	return contract.ContractTypeStandard
}

// getERC20Contract gets the appropriate ERC20 contract instance
//
//nolint:unused // This function is used in service methods
func (s *ERC20Service) getERC20Contract(contractAddr common.Address, contractTypeStr string) (ERC20Contract, error) {
	contractType := getContractType(contractTypeStr)

	switch contractType {
	case contract.ContractTypeOwnable:
		token, err := s.contractClient.GetERC20TokenOwnable(contractAddr)
		if err != nil {
			return nil, err
		}
		return &ERC20TokenOwnableWrapper{token: token}, nil
	case contract.ContractTypeStandard:
		fallthrough
	default:
		token, err := s.contractClient.GetERC20Token(contractAddr)
		if err != nil {
			return nil, err
		}
		return &ERC20TokenWrapper{token: token}, nil
	}
}

// ERC20TokenWrapper wraps ERC20Token to implement ERC20Contract interface
type ERC20TokenWrapper struct {
	token *erc20.ERC20Token
}

func (w *ERC20TokenWrapper) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	return w.token.BalanceOf(opts, account)
}

func (w *ERC20TokenWrapper) Name(opts *bind.CallOpts) (string, error) {
	return w.token.Name(opts)
}

func (w *ERC20TokenWrapper) Symbol(opts *bind.CallOpts) (string, error) {
	return w.token.Symbol(opts)
}

func (w *ERC20TokenWrapper) Decimals(opts *bind.CallOpts) (uint8, error) {
	return w.token.Decimals(opts)
}

func (w *ERC20TokenWrapper) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	return w.token.TotalSupply(opts)
}

func (w *ERC20TokenWrapper) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	return w.token.Allowance(opts, owner, spender)
}

func (w *ERC20TokenWrapper) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.Transfer(opts, to, amount)
}

func (w *ERC20TokenWrapper) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.Approve(opts, spender, amount)
}

func (w *ERC20TokenWrapper) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.TransferFrom(opts, from, to, amount)
}

func (w *ERC20TokenWrapper) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.Mint(opts, to, amount)
}

func (w *ERC20TokenWrapper) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return w.token.Burn(opts, amount)
}

func (w *ERC20TokenWrapper) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.BurnFrom(opts, from, amount)
}

// ERC20TokenOwnableWrapper wraps ERC20TokenOwnable to implement ERC20Contract interface
type ERC20TokenOwnableWrapper struct {
	token *erc20.ERC20TokenOwnable
}

func (w *ERC20TokenOwnableWrapper) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	return w.token.BalanceOf(opts, account)
}

func (w *ERC20TokenOwnableWrapper) Name(opts *bind.CallOpts) (string, error) {
	return w.token.Name(opts)
}

func (w *ERC20TokenOwnableWrapper) Symbol(opts *bind.CallOpts) (string, error) {
	return w.token.Symbol(opts)
}

func (w *ERC20TokenOwnableWrapper) Decimals(opts *bind.CallOpts) (uint8, error) {
	return w.token.Decimals(opts)
}

func (w *ERC20TokenOwnableWrapper) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	return w.token.TotalSupply(opts)
}

func (w *ERC20TokenOwnableWrapper) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	return w.token.Allowance(opts, owner, spender)
}

func (w *ERC20TokenOwnableWrapper) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.Transfer(opts, to, amount)
}

func (w *ERC20TokenOwnableWrapper) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.Approve(opts, spender, amount)
}

func (w *ERC20TokenOwnableWrapper) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.TransferFrom(opts, from, to, amount)
}

func (w *ERC20TokenOwnableWrapper) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.Mint(opts, to, amount)
}

func (w *ERC20TokenOwnableWrapper) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return w.token.Burn(opts, amount)
}

func (w *ERC20TokenOwnableWrapper) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return w.token.BurnFrom(opts, from, amount)
}
