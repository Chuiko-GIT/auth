package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Chuiko-GIT/auth/internal/config"
	"github.com/Chuiko-GIT/auth/internal/config/env"
	"github.com/Chuiko-GIT/auth/pkg/user_api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	user_api.UnimplementedUserAPIServer
	pool *pgxpool.Pool
}

func (s server) Create(ctx context.Context, req *user_api.CreateRequest) (*user_api.CreateResponse, error) {
	builderInsert := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "password_confirm", "role").
		Values(req.User.Name, req.User.Email, req.User.Password, req.User.PasswordConfirm, req.User.Role.String()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var userID int
	if err = s.pool.QueryRow(ctx, query, args...).Scan(&userID); err != nil {
		log.Fatalf("failed to select users: %v", err)
	}

	log.Printf("inserted user with id: %d", userID)

	return &user_api.CreateResponse{
		Id: int64(userID),
	}, nil
}

func (s server) Delete(ctx context.Context, req *user_api.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s server) Update(ctx context.Context, req *user_api.UpdateRequest) (*emptypb.Empty, error) {
	builderUpdate := sq.Update("users").
		PlaceholderFormat(sq.Dollar).
		Set("name", "Alex").
		Set("email", "test@gmail.com").
		Set("password", "test-test").
		Set("password_confirm", "true").
		Set("role", "admin").
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	log.Printf("updated %d rows", res.RowsAffected())

	return &emptypb.Empty{}, nil
}

func (s server) Get(ctx context.Context, req *user_api.GetRequest) (*user_api.GetResponse, error) {
	builderSelectOne := sq.Select("id", "name", "email", "password", "password_confirm", "role", "created_at", "updated_at").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var resp user_api.User
	err = s.pool.
		QueryRow(ctx, query, args...).
		Scan(resp.Id, resp.User.Name, resp.User.Email, resp.User.Password, resp.User.PasswordConfirm, resp.User.Role, resp.CreatedAt, resp.UpdatedAt)
	if err != nil {
		log.Fatalf("failed to select user: %v", err)
	}

	log.Printf("id: %d, name: %s, email: %s,password: %s,passwordConfirm: %s,role: %s, created_at: %v, updated_at: %v\n",
		resp.Id, resp.User.Name, resp.User.Email, resp.User.Password, resp.User.PasswordConfirm, resp.User.Role, resp.CreatedAt, resp.UpdatedAt)

	return &user_api.GetResponse{User: &resp}, nil

}

func main() {
	flag.Parse()
	ctx := context.Background()

	if err := config.Load(configPath); err != nil {
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
	user_api.RegisterUserAPIServer(s, server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
