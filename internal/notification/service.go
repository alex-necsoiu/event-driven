package notification

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/alex-necsoiu/event-driven/pkg/messaging"
)

// Service handles notification business logic and event consumption
type Service struct {
	subscriber messaging.Subscriber
	logger     *log.Logger
}

// NewService creates a new notification service
func NewService(subscriber messaging.Subscriber, logger *log.Logger) *Service {
	return &Service{
		subscriber: subscriber,
		logger:     logger,
	}
}

// Start starts the notification service and subscribes to events
func (s *Service) Start() error {
	// Subscribe to user events
	if err := s.subscriber.Subscribe(messaging.EventTypeUserCreated, s.handleUserCreated); err != nil {
		return fmt.Errorf("failed to subscribe to UserCreated: %w", err)
	}

	if err := s.subscriber.Subscribe(messaging.EventTypeUserUpdated, s.handleUserUpdated); err != nil {
		return fmt.Errorf("failed to subscribe to UserUpdated: %w", err)
	}

	// Subscribe to order events
	if err := s.subscriber.Subscribe(messaging.EventTypeOrderCreated, s.handleOrderCreated); err != nil {
		return fmt.Errorf("failed to subscribe to OrderCreated: %w", err)
	}

	if err := s.subscriber.Subscribe(messaging.EventTypeOrderCompleted, s.handleOrderCompleted); err != nil {
		return fmt.Errorf("failed to subscribe to OrderCompleted: %w", err)
	}

	if err := s.subscriber.Subscribe(messaging.EventTypeOrderCancelled, s.handleOrderCancelled); err != nil {
		return fmt.Errorf("failed to subscribe to OrderCancelled: %w", err)
	}

	s.logger.Println("Notification service started and subscribed to events")
	return nil
}

// Stop stops the notification service
func (s *Service) Stop() error {
	return s.subscriber.Close()
}

// handleUserCreated handles UserCreated events
func (s *Service) handleUserCreated(event messaging.Event) {
	s.logger.Printf("Handling UserCreated event: %s", event.EventType)

	var payload messaging.UserCreatedPayload
	if err := s.unmarshalPayload(event.Payload, &payload); err != nil {
		s.logger.Printf("Failed to unmarshal UserCreated payload: %v", err)
		return
	}

	// Send welcome notification
	message := fmt.Sprintf("Welcome %s! Your account has been created successfully.", payload.Name)
	s.sendNotification(payload.UserID, "welcome", message)
}

// handleUserUpdated handles UserUpdated events
func (s *Service) handleUserUpdated(event messaging.Event) {
	s.logger.Printf("Handling UserUpdated event: %s", event.EventType)

	var payload messaging.UserUpdatedPayload
	if err := s.unmarshalPayload(event.Payload, &payload); err != nil {
		s.logger.Printf("Failed to unmarshal UserUpdated payload: %v", err)
		return
	}

	// Send profile update notification
	message := fmt.Sprintf("Your profile has been updated successfully, %s.", payload.Name)
	s.sendNotification(payload.UserID, "profile_update", message)
}

// handleOrderCreated handles OrderCreated events
func (s *Service) handleOrderCreated(event messaging.Event) {
	s.logger.Printf("Handling OrderCreated event: %s", event.EventType)

	var payload messaging.OrderCreatedPayload
	if err := s.unmarshalPayload(event.Payload, &payload); err != nil {
		s.logger.Printf("Failed to unmarshal OrderCreated payload: %v", err)
		return
	}

	// Send order confirmation notification
	message := fmt.Sprintf("Your order #%s for $%.2f has been created and is being processed.", payload.OrderID, payload.Amount)
	s.sendNotification(payload.UserID, "order_confirmation", message)
}

// handleOrderCompleted handles OrderCompleted events
func (s *Service) handleOrderCompleted(event messaging.Event) {
	s.logger.Printf("Handling OrderCompleted event: %s", event.EventType)

	var payload map[string]interface{}
	if err := s.unmarshalPayload(event.Payload, &payload); err != nil {
		s.logger.Printf("Failed to unmarshal OrderCompleted payload: %v", err)
		return
	}

	userID, ok := payload["user_id"].(string)
	if !ok {
		s.logger.Printf("Invalid user_id in OrderCompleted payload")
		return
	}

	orderID, ok := payload["order_id"].(string)
	if !ok {
		s.logger.Printf("Invalid order_id in OrderCompleted payload")
		return
	}

	// Send order completion notification
	message := fmt.Sprintf("Your order #%s has been completed successfully!", orderID)
	s.sendNotification(userID, "order_completed", message)
}

// handleOrderCancelled handles OrderCancelled events
func (s *Service) handleOrderCancelled(event messaging.Event) {
	s.logger.Printf("Handling OrderCancelled event: %s", event.EventType)

	var payload map[string]interface{}
	if err := s.unmarshalPayload(event.Payload, &payload); err != nil {
		s.logger.Printf("Failed to unmarshal OrderCancelled payload: %v", err)
		return
	}

	userID, ok := payload["user_id"].(string)
	if !ok {
		s.logger.Printf("Invalid user_id in OrderCancelled payload")
		return
	}

	orderID, ok := payload["order_id"].(string)
	if !ok {
		s.logger.Printf("Invalid order_id in OrderCancelled payload")
		return
	}

	// Send order cancellation notification
	message := fmt.Sprintf("Your order #%s has been cancelled.", orderID)
	s.sendNotification(userID, "order_cancelled", message)
}

// sendNotification sends a notification to a user
func (s *Service) sendNotification(userID, notificationType, message string) {
	// In a real implementation, this would integrate with:
	// - Email service (SendGrid, AWS SES)
	// - SMS service (Twilio)
	// - Push notifications (Firebase)
	// - In-app notifications

	s.logger.Printf("Sending %s notification to user %s: %s", notificationType, userID, message)

	// Simulate notification sending
	time.Sleep(100 * time.Millisecond) // Simulate network delay

	s.logger.Printf("Notification sent successfully to user %s", userID)
}

// unmarshalPayload unmarshals event payload into the given struct
func (s *Service) unmarshalPayload(payload interface{}, target interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	return nil
}
