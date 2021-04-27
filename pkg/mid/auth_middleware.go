package mid

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/realotz/whole/pkg/token"
)

func JwtTokenWithCtx(tk *token.Token) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if info, ok := http.FromServerContext(ctx); ok {
				authorization := info.Request.Header.Get("Authorization")
				if authorization != "" {
					cc, err := tk.Decode(authorization)
					if err == nil {
						ctx = token.WithLoginContext(ctx, cc)
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
