// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"user/internal/biz"
	"user/internal/conf"
	"user/internal/data"
	"user/internal/server"
	"user/internal/service"
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
	UserRepo := data.NewUserRepo(dataData, logger)
	UserUsecase := biz.NewUserUsecase(UserRepo, logger)
	UserService := service.NewUserService(UserUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, UserService, logger)
	grpcServer := server.NewGRPCServer(confServer, UserService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
