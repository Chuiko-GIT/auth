package users

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Chuiko-GIT/auth/internal/converter"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

func (i *Implementation) Update(ctx context.Context, req *user_api.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, i.userService.Update(ctx, converter.ToUserUpdateFromDesc(req))
}
