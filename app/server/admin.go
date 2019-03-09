package server

import (
	"context"

	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

func getAdmin(ctx context.Context, store di.StoreComponent) (*record.User, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	us := store.UserStore(ctx)
	u, err := us.GetUserWithPrivate(userID)
	if err != nil {
		return nil, err
	}

	if u.Authority != int(model.Admin) {
		return nil, util.ErrUnauthenticated
	}
	return u, nil
}
