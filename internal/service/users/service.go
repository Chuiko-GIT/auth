package users

import (
	"github.com/Chuiko-GIT/auth/internal/repository"
	"github.com/Chuiko-GIT/auth/internal/service"
)

var _ service.Users = &Serv{}

type Serv struct {
	usersRepo repository.Users
}

func NewService(userRepo repository.Users) *Serv {
	return &Serv{
		usersRepo: userRepo,
	}
}
