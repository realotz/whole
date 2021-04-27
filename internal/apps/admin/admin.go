package admin

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/realotz/whole/api/admin/v1"
	"github.com/realotz/whole/internal/apps/admin/biz"
	"github.com/realotz/whole/internal/apps/admin/data"
	"github.com/realotz/whole/internal/apps/admin/service"
	"github.com/realotz/whole/internal/server"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	data.ProviderSet,
	biz.ProviderSet,
	service.NewEmployeeService,
	NewAdminApp,
)

type Admin struct {
	*server.App
}

func NewAdminApp(hs *http.Server, gs *grpc.Server, m middleware.Middleware,
	employee v1.EmployeeServiceServer,
) *Admin {
	// 用户user服务组
	hs.HandlePrefix("/admin/employee", v1.NewEmployeeServiceHandler(employee, http.Middleware(m))) // 员工服务

	// grpc
	// 用户服务
	v1.RegisterEmployeeServiceServer(gs, employee) // 员工
	return &Admin{App: server.NewApp(hs, gs)}
}
