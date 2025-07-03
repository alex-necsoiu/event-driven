package messaging

import (
	"time"
)

// EventType constants for all system events
const (
	// User events
	EventTypeUserCreated = "UserCreated"
	EventTypeUserUpdated = "UserUpdated"
	EventTypeUserDeleted = "UserDeleted"

	// Order events
	EventTypeOrderCreated   = "OrderCreated"
	EventTypeOrderUpdated   = "OrderUpdated"
	EventTypeOrderCancelled = "OrderCancelled"
	EventTypeOrderCompleted = "OrderCompleted"

	// Notification events
	EventTypeNotificationSent   = "NotificationSent"
	EventTypeNotificationFailed = "NotificationFailed"
)

// EventPayloads define the structure for different event types
type UserCreatedPayload struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type UserUpdatedPayload struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updated_at"`
}

type OrderCreatedPayload struct {
	OrderID   string  `json:"order_id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}

type OrderUpdatedPayload struct {
	OrderID   string  `json:"order_id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	UpdatedAt string  `json:"updated_at"`
}

type NotificationPayload struct {
	UserID  string `json:"user_id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	SentAt  string `json:"sent_at"`
}

// Helper functions to create events
func NewUserCreatedEvent(userID, name, email string) Event {
	return Event{
		EventType: EventTypeUserCreated,
		Payload: UserCreatedPayload{
			UserID:    userID,
			Name:      name,
			Email:     email,
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

func NewOrderCreatedEvent(orderID, userID string, amount float64) Event {
	return Event{
		EventType: EventTypeOrderCreated,
		Payload: OrderCreatedPayload{
			OrderID:   orderID,
			UserID:    userID,
			Amount:    amount,
			Status:    "pending",
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

func NewNotificationEvent(userID, notificationType, message string) Event {
	return Event{
		EventType: EventTypeNotificationSent,
		Payload: NotificationPayload{
			UserID:  userID,
			Type:    notificationType,
			Message: message,
			SentAt:  time.Now().UTC().Format(time.RFC3339),
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}
