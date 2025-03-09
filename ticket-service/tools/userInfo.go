package tools

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strings"
)

func ContextGetUserInfo(ctx context.Context, key string) string {
	var resp string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for k, v := range md {
			if k == strings.ToLower(key) {
				resp = v[0]
			}
		}
	}
	return resp
}
