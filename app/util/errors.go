package util

import (
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrUnauthenticated represents unauthenticated error
	ErrUnauthenticated = status.Error(codes.Unauthenticated, "unauthenticated")
	// ErrNotFound represents not found error
	ErrNotFound = status.Error(codes.NotFound, "not found")
)

// CodeFromHTTPStatus converts corresponding HTTP response status into the gRPC error code.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func CodeFromHTTPStatus(status int) codes.Code {
	switch status {
	case http.StatusOK:
		return codes.OK
	case http.StatusRequestTimeout:
		return codes.Canceled
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusPreconditionFailed:
		return codes.FailedPrecondition
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	}

	return codes.Internal
}

// ErrorFromRecover creates error with stacktrace from recover
func ErrorFromRecover(err interface{}) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case string:
		return errors.New(e)
	case error:
		return errors.WithStack(e)
	default:
		return errors.New("unknown panic")
	}
}
