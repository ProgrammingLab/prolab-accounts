package server

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// NewSessionServiceServer creates a new SessionServiceServer instance.
func NewSessionServiceServer(store di.StoreComponent) interface {
	api_pb.SessionServiceServer
	grapiserver.Server
} {
	return &sessionServiceServerImpl{
		StoreComponent: store,
	}
}

type sessionServiceServerImpl struct {
	di.StoreComponent
}

var errLogin = status.Error(codes.NotFound, "Invalid name or password")

func (s *sessionServiceServerImpl) GetSession(ctx context.Context, req *api_pb.GetSessionRequest) (*api_pb.Session, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *sessionServiceServerImpl) CreateSession(ctx context.Context, req *api_pb.CreateSessionRequest) (*api_pb.Session, error) {
	u, err := s.UserStore(ctx).FindUserByEmailOrName(req.GetName())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errLogin
		}
		log.Error(err)
		return nil, ErrInternalServer
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(req.GetPassword()))
	if err != nil {
		return nil, errLogin
	}

	session, err := s.SessionStore(ctx).CreateSession(model.UserID(u.ID))
	if err != nil {
		log.Error(err)
		return nil, ErrInternalServer
	}

	resp := &api_pb.Session{
		SessionId: session.ID,
	}
	return resp, nil
}

func (s *sessionServiceServerImpl) DeleteSession(ctx context.Context, req *api_pb.DeleteSessionRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
