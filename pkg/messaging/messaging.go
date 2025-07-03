package messaging

// Event represents a standard event envelope (matches event.proto)
type Event struct {
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
	Timestamp string      `json:"timestamp"`
}

// Publisher publishes events to the event bus
type Publisher interface {
	Publish(subject string, event Event) error
	Close() error
}

// Subscriber subscribes to events from the event bus
type Subscriber interface {
	Subscribe(subject string, handler func(Event)) error
	Close() error
}

// NewPublisher returns a new Publisher for the given backend (NATS, Kafka, etc.)
func NewPublisher(url string) (Publisher, error) {
	// TODO: Switch on URL/protocol to support multiple backends
	return NewNATSPublisher(url)
}

// NewSubscriber returns a new Subscriber for the given backend (NATS, Kafka, etc.)
func NewSubscriber(url string) (Subscriber, error) {
	// TODO: Switch on URL/protocol to support multiple backends
	return NewNATSSubscriber(url)
}
