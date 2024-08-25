package converter

import (
	"github.com/Chuiko-GIT/auth/internal/model"
	dbModel "github.com/Chuiko-GIT/auth/internal/repository/users/model"
)

func ToUserFromRepo(user dbModel.UserRepo) model.User {
	return model.User{
		ID:        user.ID,
		UserInfo:  ToUserInfoFromRepo(user.UserInfo),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(userInfo dbModel.UserInfoRepo) model.UserInfo {
	return model.UserInfo{
		Name:            userInfo.Name,
		Email:           userInfo.Email,
		Password:        userInfo.Password,
		PasswordConfirm: userInfo.PasswordConfirm,
		Role:            userInfo.Role,
	}
}
