package model

import (
	"database/sql"
	"time"
)

type (
	UserInfoRepo struct {
		Name            string `db:"name"`
		Email           string `db:"email"`
		Password        string `db:"password"`
		PasswordConfirm string `db:"password_confirm"`
		Role            string `db:"role"`
	}

	UserRepo struct {
		ID        int64        `db:"id"`
		UserInfo  UserInfoRepo `db:""`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}
)
