package user

import (
	"fmt"
	"log"

	"github.com/alex-necsoiu/event-driven/pkg/messaging"
)

// Service handles user business logic and event publishing
type Service struct {
	repo      Repository
	publisher messaging.Publisher
	logger    *log.Logger
}

// NewService creates a new user service
func NewService(repo Repository, publisher messaging.Publisher, logger *log.Logger) *Service {
	return &Service{
		repo:      repo,
		publisher: publisher,
		logger:    logger,
	}
}

// CreateUser creates a new user and publishes UserCreated event
func (s *Service) CreateUser(name, email string) (string, error) {
	// Create user in database
	userID, err := s.repo.CreateUser(name, email)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	s.logger.Printf("Created user: %s with ID: %s", email, userID)

	// Publish UserCreated event
	event := messaging.NewUserCreatedEvent(userID, name, email)
	if err := s.publisher.Publish(messaging.EventTypeUserCreated, event); err != nil {
		s.logger.Printf("Failed to publish UserCreated event: %v", err)
		// Don't fail the operation if event publishing fails
		// In production, you might want to retry or use outbox pattern
	}

	return userID, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(id string) (User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}
