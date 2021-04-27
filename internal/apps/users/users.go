package users

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/realotz/whole/api/users/v1"
	userV1 "github.com/realotz/whole/api/users/v1"
	"github.com/realotz/whole/internal/apps/users/biz"
	"github.com/realotz/whole/internal/apps/users/data"
	"github.com/realotz/whole/internal/apps/users/service"
	"github.com/realotz/whole/internal/conf"
	"github.com/realotz/whole/internal/server"
	"github.com/realotz/whole/pkg/token"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAuthToken,
	data.ProviderSet,
	biz.ProviderSet,
	service.NewEmployeeService,
	service.NewCustomerService,
	NewUsersApp,
)

func NewAuthToken(c *conf.UserConfig) (*token.Token, error) {
	return token.New(c.Cert.Key, c.Cert.Cert)
}

type Users struct {
	*server.App
}

func NewUsersApp(hs *http.Server, gs *grpc.Server, m middleware.Middleware,
	employee v1.EmployeeServiceServer,
	customer userV1.CustomerServiceServer,
) *Users {
	// 用户user服务组
	hs.HandlePrefix("/users/employee", v1.NewEmployeeServiceHandler(employee, http.Middleware(m)))     // 员工服务
	hs.HandlePrefix("/users/customer", userV1.NewCustomerServiceHandler(customer, http.Middleware(m))) // 客户服务

	// grpc
	// 用户服务
	v1.RegisterEmployeeServiceServer(gs, employee)     // 员工
	userV1.RegisterCustomerServiceServer(gs, customer) // 客户
	return &Users{App: server.NewApp(hs, gs)}
}
