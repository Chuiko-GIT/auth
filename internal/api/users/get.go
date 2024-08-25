package users

import (
	"context"

	"github.com/Chuiko-GIT/auth/internal/converter"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

func (i *Implementation) Get(ctx context.Context, req *user_api.GetRequest) (*user_api.GetResponse, error) {
	resp, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return &user_api.GetResponse{}, err
	}

	return &user_api.GetResponse{User: converter.ToUserFromService(resp)}, nil
}
