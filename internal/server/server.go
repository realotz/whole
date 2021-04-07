package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewMiddleware, NewHTTPServer, NewGRPCServer)

// 全局Middleware
func NewMiddleware() middleware.Middleware {
	return middleware.Chain(
		recovery.Recovery(),
		status.Server(),
		tracing.Server(),
		logging.Server(),
	)
}

type App struct {
	hs *http.Server
	gs *grpc.Server
}

func NewApp(hs *http.Server, gs *grpc.Server) *App {
	return &App{
		hs: hs,
		gs: gs,
	}
}

func (a *App) Server() []transport.Server {
	return []transport.Server{a.hs, a.gs}
}
