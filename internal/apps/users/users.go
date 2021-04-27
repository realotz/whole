package users

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	userV1 "github.com/realotz/whole/api/users/v1"
	"github.com/realotz/whole/internal/apps/users/biz"
	"github.com/realotz/whole/internal/apps/users/data"
	"github.com/realotz/whole/internal/apps/users/service"
	"github.com/realotz/whole/internal/server"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	data.ProviderSet,
	biz.ProviderSet,
	service.NewCustomerService,
	NewUsersApp,
)

type Users struct {
	*server.App
}

func NewUsersApp(hs *http.Server, gs *grpc.Server, m middleware.Middleware,
	customer userV1.CustomerServiceServer,
) *Users {
	// 用户user服务组
	hs.HandlePrefix("/users/customer", userV1.NewCustomerServiceHandler(customer, http.Middleware(m))) // 客户服务

	// grpc
	// 用户服务
	userV1.RegisterCustomerServiceServer(gs, customer) // 客户
	return &Users{App: server.NewApp(hs, gs)}
}
