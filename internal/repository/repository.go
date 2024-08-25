package repository

import (
	"context"

	"github.com/Chuiko-GIT/auth/internal/model"
)

type (
	Users interface {
		Create(ctx context.Context, user model.UserInfo) (int64, error)
		Get(ctx context.Context, id int64) (model.User, error)
		GetAll(ctx context.Context) ([]model.User, error)
		Update(ctx context.Context, req model.UpdateUser) error
		Delete(ctx context.Context, id int64) error
	}
)
