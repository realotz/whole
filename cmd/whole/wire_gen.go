// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/apps/cms/biz"
	"github.com/realotz/whole/internal/apps/cms/data"
	"github.com/realotz/whole/internal/apps/cms/service"
	biz3 "github.com/realotz/whole/internal/apps/systems/biz"
	data3 "github.com/realotz/whole/internal/apps/systems/data"
	service3 "github.com/realotz/whole/internal/apps/systems/service"
	biz2 "github.com/realotz/whole/internal/apps/users/biz"
	data2 "github.com/realotz/whole/internal/apps/users/data"
	service2 "github.com/realotz/whole/internal/apps/users/service"
	"github.com/realotz/whole/internal/conf"
	"github.com/realotz/whole/internal/server"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, error) {
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	categoryUsecase := biz.NewCategoryUsecase(categoryRepo, logger)
	categoryServiceServer := service.NewCategoryServiceService(categoryUsecase)
	messageServiceServer := service2.NewMessageService()
	data4, err := data2.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	memberRepo := data2.NewMemberRepo(data4, logger)
	memberUsecase := biz2.NewMemberUsecase(memberRepo, logger)
	memberServiceServer := service2.NewMemberService(memberUsecase)
	data5, err := data3.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	fileRepo := data3.NewFileRepo(data5, logger)
	fileUsecase := biz3.NewFileUsecase(fileRepo, logger)
	fileServiceServer := service3.NewFileServiceService(fileUsecase)
	httpServer := server.NewHTTPServer(confServer, categoryServiceServer, messageServiceServer, memberServiceServer, fileServiceServer)
	grpcServer := server.NewGRPCServer(confServer, categoryServiceServer, messageServiceServer, memberServiceServer, fileServiceServer)
	app := newApp(logger, httpServer, grpcServer)
	return app, nil
}
