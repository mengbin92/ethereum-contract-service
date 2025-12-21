// Package contract provides contract client wrapper for ERC20 token interactions.
package contract

import (
	"context"

	"eth-contract-service/internal/errors"
	"eth-contract-service/provider/contract/erc20"
	"eth-contract-service/provider/eth"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kratos/kratos/v2/log"
	pkgErrors "github.com/pkg/errors"
)

// Client wraps ERC20 contract interactions
type Client struct {
	logger *log.Helper
}

// NewClient creates a new contract client
func NewClient(logger log.Logger) *Client {
	return &Client{
		logger: log.NewHelper(logger),
	}
}

// GetERC20Token creates an ERC20Token contract instance
func (c *Client) GetERC20Token(contractAddr common.Address) (*erc20.ERC20Token, error) {
	client := eth.GetClient()
	if client == nil {
		return nil, pkgErrors.Wrap(errors.ErrClientNotInitialized, "failed to get ethereum client")
	}

	token, err := erc20.NewERC20Token(contractAddr, client)
	if err != nil {
		return nil, pkgErrors.Wrapf(err, "failed to create ERC20Token instance for address %s", contractAddr.Hex())
	}

	return token, nil
}

// CreateTransactOpts creates transaction options from private key
func (c *Client) CreateTransactOpts(ctx context.Context, privateKey []byte) (*bind.TransactOpts, error) {
	// Parse private key
	key, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return nil, pkgErrors.Wrap(errors.ErrInvalidPrivateKey, err.Error())
	}

	// Get chain ID
	chainID := eth.GetChainID()
	if chainID == nil {
		return nil, pkgErrors.Wrap(errors.ErrChainIDNotConfigured, "chain ID not configured")
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, pkgErrors.Wrap(err, "failed to create transactor")
	}

	// Set context
	auth.Context = ctx

	return auth, nil
}

// GetAddressFromPrivateKey derives the Ethereum address from a private key
func (c *Client) GetAddressFromPrivateKey(privateKey []byte) (common.Address, error) {
	key, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return common.Address{}, pkgErrors.Wrap(errors.ErrInvalidPrivateKey, err.Error())
	}
	return crypto.PubkeyToAddress(key.PublicKey), nil
}

// ValidateClient validates that the Ethereum client is initialized
func (c *Client) ValidateClient() error {
	client := eth.GetClient()
	if client == nil {
		return errors.ErrClientNotInitialized
	}
	return nil
}

