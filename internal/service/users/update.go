package users

import (
	"context"
	"errors"

	"github.com/Chuiko-GIT/auth/internal/model"
)

func (s Serv) Update(ctx context.Context, request model.UpdateUser) error {
	if err := s.usersRepo.Update(ctx, request); err != nil {
		return errors.New("failed update user")
	}
	return nil
}
