package users

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/Chuiko-GIT/auth/internal/model"
	"github.com/Chuiko-GIT/auth/internal/repository"
	"github.com/Chuiko-GIT/auth/internal/repository/users/converter"
	repoModel "github.com/Chuiko-GIT/auth/internal/repository/users/model"
)

const (
	tableUsers = "users"

	userColumnID              = "id"
	userColumnName            = "name"
	userColumnEmail           = "email"
	userColumnPassword        = "password"
	userColumnPasswordConfirm = "password_confirm"
	userColumnRole            = "role"
	userColumnCreatedAt       = "created_at"
	userColumnUpdatedAt       = "updated_at"
)

var _ repository.Users = &Repo{}

type Repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repo {
	return &Repo{db: db}
}

func (r Repo) Create(ctx context.Context, user model.UserInfo) (int64, error) {
	builderInsert := sq.Insert(tableUsers).
		PlaceholderFormat(sq.Dollar).
		Columns(userColumnName, userColumnEmail, userColumnPassword, userColumnPasswordConfirm, userColumnRole).
		Values(user.Name, user.Email, user.Password, user.PasswordConfirm, user.Role).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, errors.New(fmt.Sprintf("failed to build query: %v", err))
	}

	var userID int64
	if err = r.db.QueryRow(ctx, query, args...).Scan(&userID); err != nil {
		return 0, errors.New(fmt.Sprintf("failed to select users: %v", err))
	}

	log.Printf("inserted user with id: %d", userID)

	return userID, nil
}

func (r Repo) Get(ctx context.Context, id int64) (model.User, error) {
	builderSelectOne := sq.Select(userColumnID, userColumnName, userColumnEmail, userColumnPassword, userColumnPasswordConfirm, userColumnRole, userColumnCreatedAt, userColumnUpdatedAt).
		From(tableUsers).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{userColumnID: id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return model.User{}, errors.New(fmt.Sprintf("failed to build query: %v", err))
	}

	var resp repoModel.UserRepo
	err = r.db.
		QueryRow(ctx, query, args...).
		Scan(&resp.ID, &resp.UserInfo.Name, &resp.UserInfo.Email, &resp.UserInfo.Password, &resp.UserInfo.PasswordConfirm, &resp.UserInfo.Role, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		return model.User{}, errors.New(fmt.Sprintf("failed to select user: %v", err))
	}

	log.Printf("id: %d, name: %s, email: %s,password: %s,passwordConfirm: %s,role: %s, created_at: %v, updated_at: %v\n",
		&resp.ID, &resp.UserInfo.Name, &resp.UserInfo.Email, &resp.UserInfo.Password, &resp.UserInfo.PasswordConfirm, &resp.UserInfo.Role, &resp.CreatedAt, &resp.UpdatedAt)

	return converter.ToUserFromRepo(resp), nil
}

func (r Repo) GetAll(ctx context.Context) ([]model.User, error) {
	// TODO implement me
	panic("implement me")
}

func (r Repo) Update(ctx context.Context, request model.UpdateUser) error {
	builderUpdate := sq.Update(tableUsers).
		PlaceholderFormat(sq.Dollar).
		Set(userColumnName, request.Name).
		Set(userColumnEmail, request.Email).
		Set(userColumnUpdatedAt, time.Now()).
		Where(sq.Eq{userColumnID: request.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to build query: %v", err))
	}

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update user: %v", err))
	}

	log.Printf("updated %d rows", res.RowsAffected())

	return nil
}

func (r Repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where("id = $1", id)

	query, args, err := builder.ToSql()
	if err != nil {
		return errors.New("failed to build query")
	}

	if _, err = r.db.Exec(ctx, query, args...); err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
