package grpc

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

// Parse HTTP cookie from a gRPC server side context. It assumes that the key of the raw cookie
// text is "grpcgateway-cookie".
func GetGrpcGatewayCookie(ctx context.Context) []*http.Cookie {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	rawCookies, ok := md["grpcgateway-cookie"]
	if !ok {
		return nil
	}

	joined := strings.Join(rawCookies, "; ")
	header := http.Header{}
	header.Add("Cookie", joined)
	request := http.Request{Header: header}
	return request.Cookies()
}

// Get the first non-empty metadata value with the given key.
func GetFirstMetadataByKey(md metadata.MD, key string) (value string, ok bool) {
	for _, value := range md[key] {
		if value != "" {
			return value, true
		}
	}
	return "", false
}

// Get the first non-empty cookie value with the given key.
func GetFirstCookieByKey(cookies []*http.Cookie, key string) (value string, ok bool) {
	for _, cookie := range cookies {
		if cookie.Name == key && cookie.Value != "" {
			return cookie.Value, true
		}
	}
	return "", false
}
