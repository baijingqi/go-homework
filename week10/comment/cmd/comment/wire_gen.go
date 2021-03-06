// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"comment/internal/biz"
	"comment/internal/conf"
	"comment/internal/data"
	"comment/internal/server"
	"comment/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	commentRepo := data.NewCommentRepo(dataData, logger)
	commentUseCase := biz.NewCommentUseCase(commentRepo, logger)
	commentCountRepo := data.NewCommentCountRepo(dataData, logger)
	commentCountUseCase := biz.NewCommentCountUseCase(commentCountRepo, logger)
	commentService := service.NewCommentService(commentUseCase, commentCountUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, commentService, logger)
	grpcServer := server.NewGRPCServer(confServer, commentService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
