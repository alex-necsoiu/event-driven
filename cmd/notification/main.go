package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alex-necsoiu/event-driven/internal/notification"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	logger := log.New(os.Stdout, "[notification] ", log.LstdFlags)
	cfg := notification.LoadConfig()

	// Start event consumer (placeholder)
	go notification.StartEventConsumer(cfg, logger)

	logger.Println("Notification service started. Waiting for events...")

	// Wait for interrupt signal to gracefully shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	logger.Println("Shutting down notification service...")
	// Add cleanup logic if needed
	time.Sleep(1 * time.Second)
}
