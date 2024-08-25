package env

import (
	"errors"
	"os"

	"github.com/Chuiko-GIT/auth/internal/config"
)

var _ config.PGConfig = (*PgConfig)(nil)

const (
	dsnEnvNME = "PG_DSN"
)

type PgConfig struct {
	dsn string
}

func NewPGConfig() (*PgConfig, error) {
	dsn := os.Getenv(dsnEnvNME)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &PgConfig{
		dsn: dsn,
	}, nil
}

func (p PgConfig) DSN() string {
	return p.dsn
}
