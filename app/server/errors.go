package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrInternalServer represents internal server error
var ErrInternalServer = status.Error(codes.Internal, "Internal server error")
