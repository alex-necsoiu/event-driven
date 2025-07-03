package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alex-necsoiu/event-driven/internal/notification"
	"github.com/alex-necsoiu/event-driven/pkg/messaging"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	logger := log.New(os.Stdout, "[notification] ", log.LstdFlags)

	// Initialize messaging subscriber
	subscriber, err := messaging.NewSubscriber(os.Getenv("NATS_URL"))
	if err != nil {
		logger.Fatal("failed to create subscriber:", err)
	}
	defer subscriber.Close()

	// Initialize service
	service := notification.NewService(subscriber, logger)

	// Start the service
	if err := service.Start(); err != nil {
		logger.Fatal("failed to start service:", err)
	}

	logger.Println("Notification service started and listening for events")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logger.Println("Shutting down notification service...")
	if err := service.Stop(); err != nil {
		logger.Printf("Error stopping service: %v", err)
	}
}
