package grpc

import (
	"strings"

	"google.golang.org/grpc"
)

func MustGetMethodName(unaryRpcMethod *grpc.UnaryServerInfo) string {
	parts := strings.Split(unaryRpcMethod.FullMethod, "/")
	if len(parts) != 3 {
		panic("Invalid grpc full method name: " + unaryRpcMethod.FullMethod)
	}
	return parts[2]
}
