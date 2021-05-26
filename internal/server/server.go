package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/realotz/whole/internal/conf"
	"github.com/realotz/whole/pkg/mid"
	"github.com/realotz/whole/pkg/token"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewMiddleware, NewHTTPServer, NewGRPCServer)

// 全局Middleware
func NewMiddleware(c *conf.Data, logger log.Logger) (middleware.Middleware, error) {
	// 获取公钥实例
	jk, err := token.NewPublic(c.JwtCert)
	if err != nil {
		return nil, err
	}
	return middleware.Chain(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		// jwt tk验证为whole模式下使用，微服务模式下在网关鉴权
		mid.JwtTokenWithCtx(jk),
	), nil
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
