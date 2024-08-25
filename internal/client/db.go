package client

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type (
	Client interface {
		DB() DB
		Close() error
	}

	SQLExec interface {
		NamedExec
		QueryExec
	}

	NamedExec interface {
		ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
		ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	}

	QueryExec interface {
		ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
		QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
		QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
	}

	Ping interface {
		Ping(ctx context.Context) error
	}

	DB interface {
		SQLExec
		Ping
		Close()
	}

	Query struct {
		Name     string
		QueryRaw string
	}
)
