package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alex-necsoiu/event-driven/internal/gateway"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	_ = godotenv.Load(".env")

	// Initialize logger (placeholder)
	logger := log.New(os.Stdout, "[gateway] ", log.LstdFlags)

	// Load config (placeholder)
	cfg := gateway.LoadConfig()

	// Set up HTTP handlers
	http.HandleFunc("/users", gateway.UserHandler(cfg, logger))
	// Add more routes as needed

	logger.Println("Starting REST Gateway on port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
