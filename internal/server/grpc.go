// Package server provides server initialization for both gRPC and HTTP servers.
package server

import (
	erc20V1 "eth-contract-service/api/erc20/v1"
	"eth-contract-service/internal/conf"
	"eth-contract-service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer creates and configures a new gRPC server instance.
// It sets up middleware, network configuration, address, and timeout from the provided configuration.
// The server registers the ERC20 service implementation.
//
// Parameters:
//   - c: Server configuration containing gRPC settings
//   - logger: Logger instance for server logging
//
// Returns:
//   - *grpc.Server: A configured gRPC server ready to accept connections
func NewGRPCServer(c *conf.Server, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	// Register ERC20 service
	erc20Service := service.NewERC20Service(logger)
	erc20V1.RegisterERC20Server(srv, erc20Service)

	return srv
}

