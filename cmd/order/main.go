package main

import (
	"log"
	"net"
	"os"

	"github.com/alex-necsoiu/event-driven/internal/order"

	"github.com/alex-necsoiu/event-driven/api/proto/gen"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load(".env")
	logger := log.New(os.Stdout, "[order] ", log.LstdFlags)
	cfg := order.LoadConfig()

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		logger.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	gen.RegisterOrderServiceServer(grpcServer, order.NewOrderHandler(cfg, logger))

	logger.Println("Order gRPC service listening on", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", err)
	}
}
