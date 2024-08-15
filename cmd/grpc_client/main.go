package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Chuiko-GIT/chat/pkg/user_api"
)

const (
	address = "localhost:50051"
	userID  = 1
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	c := user_api.NewUserAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getUser, err := c.Get(ctx, &user_api.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	log.Printf(color.RedString("User get:\n"), color.GreenString("%+v", getUser.GetUser()))

	createUser, err := c.Create(ctx, &user_api.CreateRequest{})
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	log.Printf(color.RedString("User create:\n"), color.GreenString("%+v", createUser.GetId()))

	updateUser, err := c.Update(ctx, &user_api.UpdateRequest{})
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	log.Printf(color.RedString("User update:\n"), color.GreenString("%+v", updateUser.String()))

	deleteUser, err := c.Delete(ctx, &user_api.DeleteRequest{})
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}

	log.Printf(color.RedString("User delete:\n"), color.GreenString("%+v", deleteUser.String()))
}
