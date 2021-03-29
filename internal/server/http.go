package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	cmsV1 "github.com/realotz/whole/api/cms/v1"
	systemV1 "github.com/realotz/whole/api/systems/v1"
	userV1 "github.com/realotz/whole/api/users/v1"
	"github.com/realotz/whole/internal/conf"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
	category cmsV1.CategoryServiceServer,
	message userV1.MessageServiceServer,
	member userV1.MemberServiceServer,
	file systemV1.FileServiceServer,
) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	m := http.Middleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(),
		),
	)
	// cms 文章cms服务组
	srv.HandlePrefix("/cms/category", cmsV1.NewCategoryServiceHandler(category, m)) // 分类

	// 用户user服务组
	srv.HandlePrefix("/users/message", userV1.NewMessageServiceHandler(message, m)) // 消息
	srv.HandlePrefix("/users/member", userV1.NewMemberServiceHandler(member, m))    // 用户

	// 系统system服务组
	srv.HandlePrefix("/systems/file", systemV1.NewFileServiceHandler(file, m)) // 文件服务
	return srv
}
