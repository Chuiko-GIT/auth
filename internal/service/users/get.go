package users

import (
	"context"

	"github.com/Chuiko-GIT/auth/internal/model"
)

func (s Serv) Get(ctx context.Context, id int64) (model.User, error) {
	return s.usersRepo.Get(ctx, id)
}

func (s Serv) GetAll(ctx context.Context) ([]model.User, error) {
	return s.usersRepo.GetAll(ctx)
}
