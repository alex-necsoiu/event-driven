package notification

import (
	"log"
	"time"
)

// StartEventConsumer subscribes to events and logs them (placeholder)
func StartEventConsumer(cfg Config, logger *log.Logger) {
	logger.Println("Subscribing to event bus at", cfg.EventBusURL)
	// TODO: Connect to NATS/Kafka/RabbitMQ and subscribe to events
	for {
		logger.Println("[placeholder] Received event: UserCreated { ... }")
		time.Sleep(10 * time.Second)
	}
}
