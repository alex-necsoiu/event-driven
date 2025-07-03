package main

import (
	"log"
	"net"
	"os"

	"github.com/alex-necsoiu/event-driven/api/proto/gen"
	"github.com/alex-necsoiu/event-driven/internal/order"
	"github.com/alex-necsoiu/event-driven/pkg/messaging"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load(".env")
	logger := log.New(os.Stdout, "[order] ", log.LstdFlags)
	cfg := order.LoadConfig()

	// Initialize messaging publisher
	publisher, err := messaging.NewPublisher(os.Getenv("NATS_URL"))
	if err != nil {
		logger.Fatal("failed to create publisher:", err)
	}
	defer publisher.Close()

	// Initialize repository
	repo, err := order.NewPostgresRepository(os.Getenv("DATABASE_URL"), logger)
	if err != nil {
		logger.Fatal("failed to create repository:", err)
	}

	// Initialize service
	service := order.NewService(repo, publisher, logger)

	// Initialize handler
	handler := order.NewOrderHandler(service, logger)

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		logger.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	gen.RegisterOrderServiceServer(grpcServer, handler)

	logger.Println("Order gRPC service listening on", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", err)
	}
}
