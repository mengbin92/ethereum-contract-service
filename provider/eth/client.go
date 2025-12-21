// Package eth provides Ethereum client initialization and management.
// It wraps the go-ethereum client to provide a simplified interface for contract interactions.
package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"
	"time"

	"eth-contract-service/internal/conf"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var (
	// client is the global Ethereum client instance
	client *ethclient.Client
	// initClientOnce ensures the client is initialized only once (thread-safe)
	initClientOnce sync.Once
	// config stores the Ethereum configuration
	config *conf.Ethereum
	// logger stores the logger instance
	logger log.Logger
)

// Init initializes the Ethereum client connection.
// It uses sync.Once to ensure the client is initialized only once.
//
// Parameters:
//   - ctx: Context for the initialization operation
//   - cfg: Ethereum configuration containing RPC URL, chain ID, etc.
//   - logKratos: Logger instance for Ethereum client logging
//
// Returns:
//   - error: Error if initialization or connection fails
func Init(ctx context.Context, cfg *conf.Ethereum, logKratos log.Logger) error {
	if cfg == nil {
		return errors.New("ethereum config cannot be nil")
	}

	var initErr error
	initClientOnce.Do(func() {
		config = cfg
		logger = logKratos

		timeout := 30 * time.Second
		if cfg.Timeout != nil {
			timeout = cfg.Timeout.AsDuration()
		}

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		ethClient, err := ethclient.DialContext(ctx, cfg.RpcUrl)
		if err != nil {
			initErr = errors.Wrap(err, "failed to connect to Ethereum node")
			return
		}

		// Verify connection by getting chain ID
		chainID, err := ethClient.ChainID(ctx)
		if err != nil {
			initErr = errors.Wrap(err, "failed to get chain ID")
			return
		}

		expectedChainID := big.NewInt(cfg.ChainId)
		if chainID.Cmp(expectedChainID) != 0 {
			initErr = errors.Errorf("chain ID mismatch: expected %d, got %d", cfg.ChainId, chainID.Int64())
			return
		}

		client = ethClient
		log.NewHelper(logger).Infof("Ethereum client initialized: chain_id=%d, rpc_url=%s", cfg.ChainId, cfg.RpcUrl)
	})

	return initErr
}

// GetClient returns the global Ethereum client instance.
// Returns nil if the client has not been initialized.
func GetClient() *ethclient.Client {
	return client
}

// GetConfig returns the Ethereum configuration.
func GetConfig() *conf.Ethereum {
	return config
}

// GetContractAddress returns the contract address for the given contract name.
// Returns empty address if the contract is not found in the configuration.
//
// Parameters:
//   - contractName: The name of the contract (e.g., "erc20")
//
// Returns:
//   - common.Address: The contract address, or zero address if not found
func GetContractAddress(contractName string) common.Address {
	if config == nil || config.Contracts == nil {
		return common.Address{}
	}

	addrStr, ok := config.Contracts[contractName]
	if !ok || addrStr == "" {
		return common.Address{}
	}

	return common.HexToAddress(addrStr)
}

// GetChainID returns the chain ID as a big.Int.
func GetChainID() *big.Int {
	if config == nil {
		return nil
	}
	return big.NewInt(config.ChainId)
}

// NewTransactOpts creates a new transaction options for contract interactions.
// It sets the chain ID, gas price, and gas limit.
//
// Parameters:
//   - ctx: Context for the transaction
//   - from: The address that will send the transaction
//   - privateKey: The private key for signing (optional, can be nil for read-only operations)
//
// Returns:
//   - *bind.TransactOpts: Transaction options configured for the current chain
//   - error: Error if configuration fails
func NewTransactOpts(ctx context.Context, from common.Address, privateKey []byte) (*bind.TransactOpts, error) {
	if client == nil {
		return nil, errors.New("Ethereum client not initialized")
	}

	chainID := GetChainID()
	if chainID == nil {
		return nil, errors.New("chain ID not configured")
	}

	var key *ecdsa.PrivateKey
	var err error
	if privateKey != nil {
		key, err = crypto.ToECDSA(privateKey)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse private key")
		}
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	opts.From = from
	opts.Context = ctx

	// Get suggested gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get suggested gas price")
	}
	opts.GasPrice = gasPrice

	// Set gas limit (can be overridden per transaction)
	opts.GasLimit = 300000 // Default gas limit, should be adjusted based on contract

	return opts, nil
}

// NewCallOpts creates a new call options for read-only contract interactions.
//
// Parameters:
//   - ctx: Context for the call
//   - blockNumber: Block number to query (nil for latest)
//
// Returns:
//   - *bind.CallOpts: Call options configured for read-only operations
func NewCallOpts(ctx context.Context, blockNumber *big.Int) *bind.CallOpts {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNumber,
	}
	return opts
}

// CallContract performs a read-only contract call.
//
// Parameters:
//   - ctx: Context for the call
//   - contractAddr: The contract address
//   - input: The encoded function call data
//   - blockNumber: Block number to query (nil for latest)
//
// Returns:
//   - []byte: The return data from the contract call
//   - error: Error if the call fails
func CallContract(ctx context.Context, contractAddr common.Address, input []byte, blockNumber *big.Int) ([]byte, error) {
	if client == nil {
		return nil, errors.New("Ethereum client not initialized")
	}

	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: input,
	}

	result, err := client.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call contract")
	}

	return result, nil
}

// SendTransaction sends a signed transaction to the network.
//
// Parameters:
//   - ctx: Context for the transaction
//   - tx: The signed transaction
//
// Returns:
//   - error: Error if the transaction fails
func SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if client == nil {
		return errors.New("Ethereum client not initialized")
	}

	return client.SendTransaction(ctx, tx)
}

// WaitMined waits for a transaction to be mined and returns the receipt.
//
// Parameters:
//   - ctx: Context for waiting
//   - txHash: The transaction hash
//
// Returns:
//   - *types.Receipt: The transaction receipt
//   - error: Error if waiting fails or transaction reverts
func WaitMined(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	if client == nil {
		return nil, errors.New("Ethereum client not initialized")
	}

	// Get transaction by hash
	tx, isPending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction")
	}
	if isPending {
		// Wait for transaction to be mined
		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to wait for transaction")
		}

		if receipt.Status == 0 {
			return nil, errors.New("transaction reverted")
		}

		return receipt, nil
	}

	// Transaction already mined, get receipt
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction receipt")
	}

	if receipt.Status == 0 {
		return nil, errors.New("transaction reverted")
	}

	return receipt, nil
}

// PackMethod packs method parameters according to the ABI.
//
// Parameters:
//   - abi: The contract ABI
//   - method: The method name
//   - args: The method arguments
//
// Returns:
//   - []byte: The packed method call data
//   - error: Error if packing fails
func PackMethod(contractABI abi.ABI, method string, args ...interface{}) ([]byte, error) {
	return contractABI.Pack(method, args...)
}

// UnpackMethod unpacks method return data according to the ABI.
//
// Parameters:
//   - abi: The contract ABI
//   - method: The method name
//   - data: The return data
//
// Returns:
//   - []interface{}: The unpacked return values
//   - error: Error if unpacking fails
func UnpackMethod(contractABI abi.ABI, method string, data []byte) ([]interface{}, error) {
	return contractABI.Unpack(method, data)
}
