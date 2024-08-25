package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	db "github.com/Chuiko-GIT/auth/internal/client"
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
	db db.Client
}

func NewRepository(db db.Client) *Repo {
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
		return 0, err
	}

	q := db.Query{
		Name:     "users.repository.Create",
		QueryRaw: query,
	}

	var userID int64
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID); err != nil {
		return 0, err
	}

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
		return model.User{}, err
	}

	q := db.Query{
		Name:     "users.repository.Get",
		QueryRaw: query,
	}

	var resp repoModel.UserRepo
	if err = r.db.DB().ScanOneContext(ctx, &resp, q, args...); err != nil {
		return model.User{}, err
	}

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

	q := db.Query{
		Name:     "users.repository.Update",
		QueryRaw: query,
	}

	if _, err = r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (r Repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where("id = $1", id)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "users.repository.Delete",
		QueryRaw: query,
	}

	if _, err = r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return err
	}

	return nil
}
