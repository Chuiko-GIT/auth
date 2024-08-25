package main

import (
	"context"
	"log"

	"github.com/Chuiko-GIT/auth/internal/app"
)

// var (
// 	configPath string
// )
//
// func init() {
// 	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
// }

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err.Error())
	}

	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err.Error())
	}
}
