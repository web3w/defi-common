package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: add tests

func GrpcErrFromContext(ctx context.Context) error {
	switch err := ctx.Err(); err {
	case nil:
		return nil
	case context.Canceled:
		return status.Error(codes.Aborted, err.Error())
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, err.Error())
	default:
		panic("unreachable")
	}
}
