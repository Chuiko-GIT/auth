package config

import "github.com/joho/godotenv"

type (
	PGConfig interface {
		DSN() string
	}

	GRPCConfig interface {
		Address() string
	}
)

func Load(path string) error {
	return godotenv.Load(path)
}
