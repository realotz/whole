package systems

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	systemV1 "github.com/realotz/whole/api/systems/v1"
	"github.com/realotz/whole/internal/server"
	"github.com/realotz/whole/internal/services/systems/biz"
	"github.com/realotz/whole/internal/services/systems/data"
	"github.com/realotz/whole/internal/services/systems/service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	data.ProviderSet,
	biz.ProviderSet,
	service.NewFileServiceService,
	NewSystemsApp,
)

type Systems struct {
	*server.App
}

func NewSystemsApp(hs *http.Server, gs *grpc.Server, m middleware.Middleware, file systemV1.FileServiceServer) *Systems {
	// 系统system服务组
	hs.HandlePrefix("/systems/file", systemV1.NewFileServiceHandler(file, http.Middleware(m))) // 文件服务
	// grpc
	// 系统system服务组
	systemV1.RegisterFileServiceServer(gs, file) // 文件服务
	return &Systems{server.NewApp(hs, gs)}
}
