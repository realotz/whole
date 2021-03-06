// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/apps"
	"github.com/realotz/whole/internal/apps/admin"
	biz2 "github.com/realotz/whole/internal/apps/admin/biz"
	data2 "github.com/realotz/whole/internal/apps/admin/data"
	service2 "github.com/realotz/whole/internal/apps/admin/service"
	"github.com/realotz/whole/internal/apps/cms"
	biz4 "github.com/realotz/whole/internal/apps/cms/biz"
	data4 "github.com/realotz/whole/internal/apps/cms/data"
	service4 "github.com/realotz/whole/internal/apps/cms/service"
	"github.com/realotz/whole/internal/apps/systems"
	"github.com/realotz/whole/internal/apps/systems/biz"
	"github.com/realotz/whole/internal/apps/systems/data"
	"github.com/realotz/whole/internal/apps/systems/service"
	"github.com/realotz/whole/internal/apps/users"
	biz3 "github.com/realotz/whole/internal/apps/users/biz"
	data3 "github.com/realotz/whole/internal/apps/users/data"
	service3 "github.com/realotz/whole/internal/apps/users/service"
	"github.com/realotz/whole/internal/conf"
	"github.com/realotz/whole/internal/server"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, userConfig *conf.UserConfig, logger log.Logger) (*kratos.App, error) {
	httpServer := server.NewHTTPServer(confServer)
	middleware, err := server.NewMiddleware(confData, logger)
	if err != nil {
		return nil, err
	}
	grpcServer := server.NewGRPCServer(confServer, middleware)
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	fileRepo := data.NewFileRepo(dataData, logger)
	fileUsecase := biz.NewFileUsecase(fileRepo, logger)
	fileServiceServer := service.NewFileServiceService(fileUsecase)
	systemsSystems := systems.NewSystemsApp(httpServer, grpcServer, middleware, fileServiceServer)
	data5, err := data2.NewData(confData, logger, userConfig)
	if err != nil {
		return nil, err
	}
	employeeRepo := data2.NewEmployeeRepo(data5, logger)
	employeeUsecase := biz2.NewEmployeeUsecase(employeeRepo, logger)
	employeeServiceServer := service2.NewEmployeeService(employeeUsecase, confData)
	adminAdmin := admin.NewAdminApp(httpServer, grpcServer, middleware, employeeServiceServer)
	data6, err := data3.NewData(confData, logger, userConfig)
	if err != nil {
		return nil, err
	}
	customerRepo := data3.NewCustomerRepo(data6, logger)
	customerUsecase := biz3.NewCustomerUsecase(customerRepo, logger)
	customerServiceServer := service3.NewCustomerService(customerUsecase)
	usersUsers := users.NewUsersApp(httpServer, grpcServer, middleware, customerServiceServer)
	data7, err := data4.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	categoryRepo := data4.NewCategoryRepo(data7, logger)
	categoryUsecase := biz4.NewCategoryUsecase(categoryRepo, logger)
	categoryServiceServer := service4.NewCategoryService(categoryUsecase)
	cmsCms := cms.NewCmsApp(httpServer, grpcServer, middleware, categoryServiceServer)
	app := apps.NewApps(systemsSystems, adminAdmin, usersUsers, cmsCms)
	kratosApp := newApp(logger, app)
	return kratosApp, nil
}
