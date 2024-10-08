package app

import (
	"context"
	"log"

	"github.com/Chuiko-GIT/auth/internal/api/users"
	uImpl "github.com/Chuiko-GIT/auth/internal/api/users"
	db "github.com/Chuiko-GIT/auth/internal/client"
	"github.com/Chuiko-GIT/auth/internal/client/db/pg"
	"github.com/Chuiko-GIT/auth/internal/closer"
	"github.com/Chuiko-GIT/auth/internal/config"
	"github.com/Chuiko-GIT/auth/internal/config/env"
	"github.com/Chuiko-GIT/auth/internal/repository"
	uRepo "github.com/Chuiko-GIT/auth/internal/repository/users"
	"github.com/Chuiko-GIT/auth/internal/service"
	uServise "github.com/Chuiko-GIT/auth/internal/service/users"
)

type ServiceProvider struct {
	pgConfig        config.PGConfig
	grpcConfig      config.GRPCConfig
	dbClient        db.Client
	usersRepository repository.Users
	usersService    service.Users
	usersImpl       *users.Implementation
}

func newServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *ServiceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		if err = cl.DB().Ping(ctx); err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *ServiceProvider) UsersRepository(ctx context.Context) repository.Users {
	if s.usersRepository == nil {
		s.usersRepository = uRepo.NewRepository(s.DBClient(ctx))
	}

	return s.usersRepository
}

func (s *ServiceProvider) UsersService(ctx context.Context) service.Users {
	if s.usersService == nil {
		s.usersService = uServise.NewService(s.UsersRepository(ctx))
	}

	return s.usersService
}

func (s *ServiceProvider) UsersImpl(ctx context.Context) *users.Implementation {
	if s.usersImpl == nil {
		s.usersImpl = uImpl.NewImplementation(s.UsersService(ctx))
	}

	return s.usersImpl
}
