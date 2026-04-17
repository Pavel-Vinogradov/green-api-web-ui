package main

import (
	"fmt"
	"green-api-web-ui/cmd/cli"
	"green-api-web-ui/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.LoadConfig()

	app, err := cli.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	fiberApp := app.RunApi()

	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	go func() {
		if err := fiberApp.Listen(addr); err != nil {
			log.Fatalf("server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("shutdown signal received")

	if err := fiberApp.Shutdown(); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
