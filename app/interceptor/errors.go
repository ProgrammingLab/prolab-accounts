package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type statusError interface {
	Error() string
	GRPCStatus() *status.Status
}

// ErrInternalServer represents internal server error
var errInternalServer = status.Error(codes.Internal, "internal server error")

// ErrorUnaryServerInterceptor returns the error handling interceptor
func ErrorUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err == nil {
			return
		}

		st, ok := err.(statusError)
		if ok {
			grpclog.Error(st)
			return
		}
		grpclog.Errorf("%+v", err)
		return resp, errInternalServer
	}
}
