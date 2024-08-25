package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	usersAPI "github.com/Chuiko-GIT/auth/internal/api/users"
	"github.com/Chuiko-GIT/auth/internal/config"
	"github.com/Chuiko-GIT/auth/internal/config/env"
	"github.com/Chuiko-GIT/auth/internal/repository/users"
	srv "github.com/Chuiko-GIT/auth/internal/service/users"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	if err := config.Load("local.env"); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	repoUsers := users.NewRepository(pool)
	serviceUser := srv.NewService(repoUsers)

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_api.RegisterUserAPIServer(s, usersAPI.NewImplementation(serviceUser))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
