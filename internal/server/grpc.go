package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	cmsV1 "github.com/realotz/whole/api/cms/v1"
	systemV1 "github.com/realotz/whole/api/systems/v1"
	userV1 "github.com/realotz/whole/api/users/v1"
	"github.com/realotz/whole/internal/conf"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server,
	category cmsV1.CategoryServiceServer,
	message userV1.MessageServiceServer,
	member userV1.MemberServiceServer,
	file systemV1.FileServiceServer,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				status.Server(),
				tracing.Server(),
				logging.Server(),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	// cms 文章cms服务组
	cmsV1.RegisterCategoryServiceServer(srv, category) //分类

	// 用户user服务组
	userV1.RegisterMessageServiceServer(srv, message) // 消息
	userV1.RegisterMemberServiceServer(srv, member)   // 用户

	// 系统system服务组
	systemV1.RegisterFileServiceServer(srv, file) // 文件服务
	return srv
}
