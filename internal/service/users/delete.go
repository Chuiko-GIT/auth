package users

import (
	"context"
	"errors"
)

func (s Serv) Delete(ctx context.Context, id int64) error {
	if err := s.usersRepo.Delete(ctx, id); err != nil {
		return errors.New("failed delete user")
	}
	return nil
}
