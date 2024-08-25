package model

import (
	"database/sql"
	"time"
)

type (
	UserInfo struct {
		Name            string
		Email           string
		Password        string
		PasswordConfirm string
		Role            string
	}

	User struct {
		ID        int64
		UserInfo  UserInfo
		CreatedAt time.Time
		UpdatedAt sql.NullTime
	}

	UpdateUser struct {
		ID    int64
		Name  string
		Email string
	}
)
