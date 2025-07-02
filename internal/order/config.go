package order

import "os"

type Config struct {
	GRPCPort    string
	DBConn      string
	EventBusURL string
}

func LoadConfig() Config {
	return Config{
		GRPCPort:    getEnv("ORDER_GRPC_PORT", "50052"),
		DBConn:      getEnv("ORDER_DB_CONN", "postgres://user:pass@localhost:5432/orders?sslmode=disable"),
		EventBusURL: getEnv("EVENT_BUS_URL", "nats://localhost:4222"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
