package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Chuiko-GIT/auth/internal/model"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

func ToUserFromService(user model.User) *user_api.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &user_api.User{
		Id:        user.ID,
		User:      ToUserInfoFromService(user.UserInfo),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(userInfo model.UserInfo) *user_api.UserInfo {
	return &user_api.UserInfo{
		Name:            userInfo.Name,
		Email:           userInfo.Email,
		Password:        userInfo.Password,
		PasswordConfirm: userInfo.PasswordConfirm,
		Role:            user_api.Role(user_api.Role_value[userInfo.Role]),
	}
}

func ToUserInfoFromDesc(userInfo *user_api.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:            userInfo.Name,
		Email:           userInfo.Name,
		Password:        userInfo.Password,
		PasswordConfirm: userInfo.PasswordConfirm,
		Role:            userInfo.Role.String(),
	}
}

func ToUserUpdateFromDesc(userUpdate *user_api.UpdateRequest) model.UpdateUser {
	return model.UpdateUser{
		ID:    userUpdate.Id,
		Name:  userUpdate.User.Name.Value,
		Email: userUpdate.User.Email.Value,
	}
}
