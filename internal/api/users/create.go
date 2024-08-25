package users

import (
	"context"

	"github.com/Chuiko-GIT/auth/internal/converter"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

func (i *Implementation) Create(ctx context.Context, req *user_api.CreateRequest) (*user_api.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetUser()))
	if err != nil {
		return &user_api.CreateResponse{}, err
	}
	return &user_api.CreateResponse{Id: id}, nil
}
