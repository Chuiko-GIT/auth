package users

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

func (i *Implementation) Delete(ctx context.Context, req *user_api.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, i.userService.Delete(ctx, req.Id)
}
