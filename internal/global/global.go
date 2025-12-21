// Package global provides global variables and initialization functions for the application.
// It manages the logger and other global resources that are used throughout the application.
package global

import (
	"context"

	"eth-contract-service/internal/conf"
	"eth-contract-service/provider/cache"
	"eth-contract-service/provider/db"
	"eth-contract-service/provider/eth"
	"eth-contract-service/provider/keystore"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// Logger is the global logger instance used throughout the application.
	Logger *log.Helper
)

// Init initializes global variables including the logger.
// It initializes database, cache, and object storage connections based on the bootstrap configuration.
//
// Parameters:
//   - bc: The bootstrap configuration containing log, data, and other settings
//   - logger: The logger instance to use for application logging
//
// The function will panic if critical initialization steps fail:
//   - Bootstrap configuration is nil
//   - Database initialization fails
func Init(bc *conf.Bootstrap, logger log.Logger) {
	if bc == nil {
		panic("bootstrap config cannot be nil")
	}

	Logger = log.NewHelper(logger)
	Logger.Infof("logger initialized: %v", bc.Log)

	Logger.Infof("database initialized")
	err := db.Init(context.Background(), bc.Data.Database, logger)
	if err != nil {
		panic(err)
	}

	err = cache.InitRedis(context.Background(), bc.Data.Redis, logger)
	if err != nil {
		Logger.Warnf("redis initialization failed: %v", err)
	}

	// Initialize Ethereum client if configured
	if bc.Ethereum != nil {
		err = eth.Init(context.Background(), bc.Ethereum, logger)
		if err != nil {
			Logger.Warnf("ethereum client initialization failed: %v", err)
		} else {
			Logger.Infof("ethereum client initialized")
		}
	} else {
		Logger.Warnf("ethereum configuration not found, skipping initialization")
	}

	// Initialize admin keystore if configured
	if bc.GetAdmin() != nil {
		err = keystore.Init(context.Background(), bc.GetAdmin(), logger)
		if err != nil {
			Logger.Warnf("admin keystore initialization failed: %v", err)
		} else {
			Logger.Infof("admin keystore initialized: address=%s", keystore.GetAdminAddress().Hex())
		}
	} else {
		Logger.Warnf("admin configuration not found, skipping keystore initialization")
	}
}
