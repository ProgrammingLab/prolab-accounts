package util

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrInternalServer represents internal server error
	ErrInternalServer = status.Error(codes.Internal, "Internal server error")
	// ErrUnauthenticated represents unauthenticated error
	ErrUnauthenticated = status.Error(codes.Unauthenticated, "Unauthenticated")
)
