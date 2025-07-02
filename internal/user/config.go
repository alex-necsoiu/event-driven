package user

import "os"

type Config struct {
	GRPCPort    string
	DBConn      string
	EventBusURL string
}

func LoadConfig() Config {
	return Config{
		GRPCPort:    getEnv("USER_GRPC_PORT", "50051"),
		DBConn:      getEnv("USER_DB_CONN", "postgres://user:pass@localhost:5432/users?sslmode=disable"),
		EventBusURL: getEnv("EVENT_BUS_URL", "nats://localhost:4222"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
