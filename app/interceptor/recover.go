package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
)

// RecoverUnaryServerInterceptor returns the recover interceptor
func RecoverUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if e := util.ErrorFromRecover(recover()); e != nil {
				err = errInternalServer
				grpclog.Errorf("recover: %+v", e)
			}
		}()
		return handler(ctx, req)
	}
}
