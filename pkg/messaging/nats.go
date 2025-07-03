package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// NATSPublisher implements Publisher for NATS
type NATSPublisher struct {
	conn *nats.Conn
}

func NewNATSPublisher(url string) (*NATSPublisher, error) {
	opts := []nats.Option{
		nats.Name("event-driven-publisher"),
		nats.ReconnectWait(time.Second),
		nats.MaxReconnects(10),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("NATS Publisher disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS Publisher reconnected to %s", nc.ConnectedUrl())
		}),
	}

	conn, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return &NATSPublisher{conn: conn}, nil
}

func (p *NATSPublisher) Publish(subject string, event Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	if err := p.conn.Publish(subject, data); err != nil {
		return fmt.Errorf("failed to publish event to %s: %w", subject, err)
	}

	log.Printf("Published event %s to subject %s", event.EventType, subject)
	return nil
}

func (p *NATSPublisher) Close() error {
	if p.conn != nil {
		p.conn.Close()
	}
	return nil
}

// NATSSubscriber implements Subscriber for NATS
type NATSSubscriber struct {
	conn *nats.Conn
	subs []*nats.Subscription
}

func NewNATSSubscriber(url string) (*NATSSubscriber, error) {
	opts := []nats.Option{
		nats.Name("event-driven-subscriber"),
		nats.ReconnectWait(time.Second),
		nats.MaxReconnects(10),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("NATS Subscriber disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS Subscriber reconnected to %s", nc.ConnectedUrl())
		}),
	}

	conn, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return &NATSSubscriber{
		conn: conn,
		subs: make([]*nats.Subscription, 0),
	}, nil
}

func (s *NATSSubscriber) Subscribe(subject string, handler func(Event)) error {
	sub, err := s.conn.Subscribe(subject, func(msg *nats.Msg) {
		var event Event
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Failed to unmarshal event from %s: %v", subject, err)
			return
		}

		log.Printf("Received event %s from subject %s", event.EventType, subject)
		handler(event)
	})

	if err != nil {
		return fmt.Errorf("failed to subscribe to %s: %w", subject, err)
	}

	s.subs = append(s.subs, sub)
	log.Printf("Subscribed to subject: %s", subject)
	return nil
}

func (s *NATSSubscriber) Close() error {
	for _, sub := range s.subs {
		sub.Unsubscribe()
	}
	if s.conn != nil {
		s.conn.Close()
	}
	return nil
}

// Example usage (in service):
// pub, _ := messaging.NewPublisher(natsURL)
// pub.Publish("UserCreated", messaging.Event{...})
// sub, _ := messaging.NewSubscriber(natsURL)
// sub.Subscribe("UserCreated", func(e messaging.Event) { ... })
