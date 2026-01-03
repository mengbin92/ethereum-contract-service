// Package server provides server initialization for both gRPC and HTTP servers.
package server

import (
	erc20V1 "eth-contract-service/api/erc20/v1"
	erc1155V1 "eth-contract-service/api/erc1155/v1"
	erc721V1 "eth-contract-service/api/erc721/v1"
	"eth-contract-service/internal/conf"
	"eth-contract-service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer creates and configures a new HTTP server instance.
// It sets up middleware, network configuration, address, and timeout from the provided configuration.
// The server registers the ERC20 service HTTP handler.
//
// Note: CORS is handled by Nginx proxy, so we don't configure it here.
// This allows the backend service to focus on business logic.
//
// Parameters:
//   - c: Server configuration containing HTTP settings
//   - logger: Logger instance for server logging
//
// Returns:
//   - *http.Server: A configured HTTP server ready to accept connections
func NewHTTPServer(c *conf.Server, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// Register health check endpoint
	srv.Route("/").GET("/health", func(ctx http.Context) error {
		return ctx.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	// Register ERC20 service
	erc20Service := service.NewERC20Service(logger)
	erc20V1.RegisterERC20HTTPServer(srv, erc20Service)

	// Register ERC1155 service
	erc1155Service := service.NewERC1155Service(logger)
	erc1155V1.RegisterERC1155HTTPServer(srv, erc1155Service)

	// Register ERC721 service
	erc721Service := service.NewERC721Service(logger)
	erc721V1.RegisterERC721HTTPServer(srv, erc721Service)

	return srv
}
