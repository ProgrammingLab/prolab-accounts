package interceptor

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

const (
	// AuthorizationKey is metadata key
	AuthorizationKey = "authorization"
	// SessionAuthorizationType is the type in authorization header
	SessionAuthorizationType = "session"
)

var (
	// ErrMetadataNotFound is returned when metadata not found in context
	ErrMetadataNotFound = errors.New("metadata not found in context")
	// ErrInvalidAuthorizationMetadata is returned when authorization metadata is invalid
	ErrInvalidAuthorizationMetadata = status.Error(codes.InvalidArgument, "Invalid authorization metadata")
)

type currentUserIDKey struct{}
type sessionIDKey struct{}

// GetCurrentUserID returns the current user's id from context
func GetCurrentUserID(ctx context.Context) (id model.UserID, ok bool) {
	v := ctx.Value(currentUserIDKey{})
	id, ok = v.(model.UserID)
	return
}

// GetSessionID returns the session id from context
func GetSessionID(ctx context.Context) (id string, ok bool) {
	v := ctx.Value(sessionIDKey{})
	id, ok = v.(string)
	return
}

// Authorizator provide the authorization interceptor
type Authorizator struct {
	di.StoreComponent
}

// NewAuthorizator returns new Authorizator
func NewAuthorizator(store di.StoreComponent) *Authorizator {
	return &Authorizator{
		StoreComponent: store,
	}
}

// UnaryServerInterceptor returns authorization UnaryServerInterceptor
func (a *Authorizator) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return a.authorization
}

func (a *Authorizator) authorization(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	d, err := fromMeta(ctx, AuthorizationKey)
	if err != nil {
		grpclog.Error(err)
		return handler(ctx, req)
	}

	if strings.Index(d, SessionAuthorizationType) != 0 {
		return nil, ErrInvalidAuthorizationMetadata
	}

	sessionID := strings.TrimSpace(d[len(SessionAuthorizationType):])
	s, err := a.SessionStore(ctx).GetSession(sessionID)
	if err != nil {
		grpclog.Error(err)
		return nil, util.ErrUnauthenticated
	}

	newCtx := context.WithValue(ctx, currentUserIDKey{}, s.UserID)
	newCtx = context.WithValue(newCtx, sessionIDKey{}, sessionID)
	return handler(newCtx, req)
}

func fromMeta(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.Wrap(ErrMetadataNotFound, "failed to get meradata from incoming contex")
	}

	h := md.Get(key)
	if len(h) == 0 {
		return "", errors.Errorf("metadata not found: key = %s", key)
	}
	return h[0], nil
}
