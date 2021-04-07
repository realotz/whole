package token

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/metadata"
)

func JwtWithAuthInfo(jt *Token) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var token string
			if info, ok := http.FromServerContext(ctx); ok {
				token = info.Request.Header.Get("Authorization")
			} else if _, ok := grpc.FromServerContext(ctx); ok {
				if md, ok := metadata.FromIncomingContext(ctx); ok {
					mdtk := md.Get("Authorization")
					if len(mdtk) > 0 {
						token = md.Get("Authorization")[0]
					}
				}
			}
			if token != "" {
				cc, err := jt.Decode(token)
				if err == nil {
					ctx = WithLoginContext(ctx, cc)
				}
			}
			return handler(ctx, req)
		}
	}
}
