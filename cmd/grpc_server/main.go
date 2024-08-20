package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Chuiko-GIT/chat/pkg/user_api"
)

const (
	grpcPort = 50051
)

type server struct {
	user_api.UnimplementedUserAPIServer
}

func (s server) Create(ctx context.Context, req *user_api.CreateRequest) (*user_api.CreateResponse, error) {
	return &user_api.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s server) Delete(ctx context.Context, req *user_api.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s server) Update(ctx context.Context, req *user_api.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s server) Get(ctx context.Context, req *user_api.GetRequest) (*user_api.GetResponse, error) {
	return &user_api.GetResponse{
		User: &user_api.User{
			Id: req.GetId(),
			User: &user_api.UserInfo{
				Name:            "TEST USER",
				Email:           gofakeit.Email(),
				Password:        gofakeit.Password(true, true, true, true, true, 1),
				PasswordConfirm: gofakeit.BeerName(),
				Role:            user_api.Role_USER,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_api.RegisterUserAPIServer(s, server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
