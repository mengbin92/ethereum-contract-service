// Package keystore provides keystore v3 file loading and management.
package keystore

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"sync"

	"eth-contract-service/internal/conf"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kratos/kratos/v2/log"
	pkgErrors "github.com/pkg/errors"
)

var (
	// adminKey is the loaded admin private key
	adminKey *ecdsa.PrivateKey
	// adminAddress is the admin address derived from the private key
	adminAddress common.Address
	// initOnce ensures the keystore is loaded only once
	initOnce sync.Once
	// initErr stores any error during initialization
	initErr error
)

// Init loads the admin keystore from the provided configuration.
// It uses sync.Once to ensure the keystore is loaded only once.
//
// Parameters:
//   - ctx: Context for the initialization operation
//   - cfg: Admin configuration containing keystore path and password
//   - logger: Logger instance for keystore logging
//
// Returns:
//   - error: Error if loading or decryption fails
func Init(ctx context.Context, cfg *conf.Admin, logger log.Logger) error {
	if cfg == nil {
		return pkgErrors.New("admin config cannot be nil")
	}

	initOnce.Do(func() {
		if cfg.KeystorePath == "" {
			initErr = pkgErrors.New("keystore_path cannot be empty")
			return
		}

		if cfg.KeystorePassword == "" {
			initErr = pkgErrors.New("keystore_password cannot be empty")
			return
		}

		// Read keystore file
		keyJSON, err := ioutil.ReadFile(cfg.KeystorePath)
		if err != nil {
			initErr = pkgErrors.Wrapf(err, "failed to read keystore file: %s", cfg.KeystorePath)
			return
		}

		// Decrypt keystore
		key, err := keystore.DecryptKey(keyJSON, cfg.KeystorePassword)
		if err != nil {
			initErr = pkgErrors.Wrap(err, "failed to decrypt keystore")
			return
		}

		adminKey = key.PrivateKey
		adminAddress = crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

		// Verify address if provided
		if cfg.Address != "" {
			expectedAddr := common.HexToAddress(cfg.Address)
			if adminAddress != expectedAddr {
				initErr = pkgErrors.Errorf("keystore address mismatch: expected %s, got %s", expectedAddr.Hex(), adminAddress.Hex())
				return
			}
		}

		log.NewHelper(logger).Infof("admin keystore loaded: address=%s, path=%s", adminAddress.Hex(), cfg.KeystorePath)
	})

	return initErr
}

// GetAdminKey returns the admin private key.
// Returns nil if the keystore has not been initialized.
func GetAdminKey() *ecdsa.PrivateKey {
	return adminKey
}

// GetAdminAddress returns the admin address.
// Returns zero address if the keystore has not been initialized.
func GetAdminAddress() common.Address {
	return adminAddress
}

// IsInitialized returns true if the keystore has been initialized.
func IsInitialized() bool {
	return adminKey != nil
}

