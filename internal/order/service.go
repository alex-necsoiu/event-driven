package order

import (
	"fmt"
	"log"
	"time"

	"github.com/alex-necsoiu/event-driven/pkg/messaging"
)

// Service handles order business logic and event publishing
type Service struct {
	repo      Repository
	publisher messaging.Publisher
	logger    *log.Logger
}

// NewService creates a new order service
func NewService(repo Repository, publisher messaging.Publisher, logger *log.Logger) *Service {
	return &Service{
		repo:      repo,
		publisher: publisher,
		logger:    logger,
	}
}

// CreateOrder creates a new order and publishes OrderCreated event
func (s *Service) CreateOrder(userID string, amount float64) (string, error) {
	// Create order in database
	orderID, err := s.repo.CreateOrder(userID, amount)
	if err != nil {
		return "", fmt.Errorf("failed to create order: %w", err)
	}

	s.logger.Printf("Created order: %s for user: %s, amount: %.2f", orderID, userID, amount)

	// Publish OrderCreated event
	event := messaging.NewOrderCreatedEvent(orderID, userID, amount)
	if err := s.publisher.Publish(messaging.EventTypeOrderCreated, event); err != nil {
		s.logger.Printf("Failed to publish OrderCreated event: %v", err)
		// Don't fail the operation if event publishing fails
	}

	return orderID, nil
}

// GetOrder retrieves an order by ID
func (s *Service) GetOrder(id string) (Order, error) {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		return Order{}, fmt.Errorf("failed to get order: %w", err)
	}
	return order, nil
}

// UpdateOrderStatus updates order status and publishes OrderUpdated event
func (s *Service) UpdateOrderStatus(orderID, status string) error {
	order, err := s.repo.GetOrder(orderID)
	if err != nil {
		return fmt.Errorf("failed to get order for status update: %w", err)
	}

	// Update order status
	if err := s.repo.UpdateOrderStatus(orderID, status); err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	s.logger.Printf("Updated order %s status to: %s", orderID, status)

	// Publish OrderUpdated event
	event := messaging.Event{
		EventType: messaging.EventTypeOrderUpdated,
		Payload: messaging.OrderUpdatedPayload{
			OrderID:   orderID,
			UserID:    order.UserID,
			Amount:    order.Amount,
			Status:    status,
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	if err := s.publisher.Publish(messaging.EventTypeOrderUpdated, event); err != nil {
		s.logger.Printf("Failed to publish OrderUpdated event: %v", err)
	}

	// Publish specific status events
	switch status {
	case "cancelled":
		s.publishOrderCancelledEvent(orderID, order.UserID)
	case "completed":
		s.publishOrderCompletedEvent(orderID, order.UserID)
	}

	return nil
}

// publishOrderCancelledEvent publishes OrderCancelled event
func (s *Service) publishOrderCancelledEvent(orderID, userID string) {
	event := messaging.Event{
		EventType: messaging.EventTypeOrderCancelled,
		Payload: map[string]interface{}{
			"order_id":     orderID,
			"user_id":      userID,
			"cancelled_at": time.Now().UTC().Format(time.RFC3339),
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	if err := s.publisher.Publish(messaging.EventTypeOrderCancelled, event); err != nil {
		s.logger.Printf("Failed to publish OrderCancelled event: %v", err)
	}
}

// publishOrderCompletedEvent publishes OrderCompleted event
func (s *Service) publishOrderCompletedEvent(orderID, userID string) {
	event := messaging.Event{
		EventType: messaging.EventTypeOrderCompleted,
		Payload: map[string]interface{}{
			"order_id":     orderID,
			"user_id":      userID,
			"completed_at": time.Now().UTC().Format(time.RFC3339),
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	if err := s.publisher.Publish(messaging.EventTypeOrderCompleted, event); err != nil {
		s.logger.Printf("Failed to publish OrderCompleted event: %v", err)
	}
}
