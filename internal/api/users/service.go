package users

import (
	"github.com/Chuiko-GIT/auth/internal/service"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

type Implementation struct {
	user_api.UnimplementedUserAPIServer
	userService service.Users
}

func NewImplementation(userService service.Users) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
