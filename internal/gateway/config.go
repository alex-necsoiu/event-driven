package gateway

import (
	"os"
)

type Config struct {
	Port          string
	UserGRPCAddr  string
	OrderGRPCAddr string
}

// LoadConfig loads config from env or defaults
func LoadConfig() Config {
	return Config{
		Port:          getEnv("GATEWAY_PORT", "8080"),
		UserGRPCAddr:  getEnv("USER_GRPC_ADDR", "localhost:50051"),
		OrderGRPCAddr: getEnv("ORDER_GRPC_ADDR", "localhost:50052"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
