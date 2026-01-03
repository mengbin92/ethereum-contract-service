// Package server provides server initialization for both gRPC and HTTP servers.
package server

import (
	erc1155V1 "eth-contract-service/api/erc1155/v1"
	erc20V1 "eth-contract-service/api/erc20/v1"
	erc721V1 "eth-contract-service/api/erc721/v1"
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

	// Register ERC1155 service
	erc1155Service := service.NewERC1155Service(logger)
	erc1155V1.RegisterERC1155Server(srv, erc1155Service)

	// Register ERC721 service
	erc721Service := service.NewERC721Service(logger)
	erc721V1.RegisterERC721Server(srv, erc721Service)

	return srv
}
