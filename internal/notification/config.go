package notification

import "os"

type Config struct {
	EventBusURL string
}

func LoadConfig() Config {
	return Config{
		EventBusURL: getEnv("EVENT_BUS_URL", "nats://localhost:4222"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
