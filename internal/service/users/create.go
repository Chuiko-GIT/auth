package users

import (
	"context"
	"errors"

	"github.com/Chuiko-GIT/auth/internal/model"
)

func (s Serv) Create(ctx context.Context, user model.UserInfo) (int64, error) {
	id, err := s.usersRepo.Create(ctx, user)
	if err != nil {
		return 0, errors.New("failed create user")
	}
	return id, nil
}
