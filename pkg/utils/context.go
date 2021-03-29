package utils

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/metadata"
	"net"
	"strings"
)

func GetClientIp(ctx context.Context) string {
	if info, ok := http.FromServerContext(ctx); ok {
		clientIP := info.Request.Header.Get("X-Forwarded-For")
		clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
		if clientIP == "" {
			clientIP = strings.TrimSpace(info.Request.Header.Get("X-Real-Ip"))
		}
		if clientIP != "" {
			return clientIP
		}
		if addr := info.Request.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
			return addr
		}
		if ip, _, err := net.SplitHostPort(strings.TrimSpace(info.Request.RemoteAddr)); err == nil {
			return ip
		}
	} else if _, ok = grpc.FromServerContext(ctx); ok {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ""
		}
		rips := md.Get("x-real-ip")
		if len(rips) == 0 {
			return ""
		}
		return rips[0]
	}
	return ""
}
