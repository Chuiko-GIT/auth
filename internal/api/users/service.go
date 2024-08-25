package users

import (
	"github.com/Chuiko-GIT/auth/internal/service"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

// В курсе решили назвать название этого файла тоже service_provider.go
// В курсе так же не придумали название данному слою, поэтому Implementation

type Implementation struct {
	user_api.UnimplementedUserAPIServer
	userService service.Users
}

func NewImplementation(userService service.Users) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
