package token

import "context"

type LoginInfoKey struct{}

func WithLoginContext(ctx context.Context, loginInfo *CustomClaims) context.Context {
	return context.WithValue(ctx, LoginInfoKey{}, loginInfo)
}

func FormLoginContext(ctx context.Context) *CustomClaims {
	data := ctx.Value(LoginInfoKey{})
	if data != nil {
		return data.(*CustomClaims)
	}
	return nil
}
