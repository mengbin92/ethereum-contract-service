package main

import (
	"eth-contract-service/internal/conf"
	"eth-contract-service/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(confServer, logger)
	httpServer := server.NewHTTPServer(confServer, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, nil, nil
}

