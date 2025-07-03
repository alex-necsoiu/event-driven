package user

import (
	"context"
	"log"

	"github.com/alex-necsoiu/event-driven/api/proto/gen"
)

// UserHandler implements the gRPC UserServiceServer
type UserHandler struct {
	gen.UnimplementedUserServiceServer
	service *Service
	logger  *log.Logger
}

func NewUserHandler(service *Service, logger *log.Logger) *UserHandler {
	return &UserHandler{service: service, logger: logger}
}

// CreateUser handles user creation and publishes an event
func (h *UserHandler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.UserResponse, error) {
	h.logger.Printf("CreateUser called for: %s (%s)", req.Name, req.Email)

	userID, err := h.service.CreateUser(req.Name, req.Email)
	if err != nil {
		h.logger.Printf("Failed to create user: %v", err)
		return &gen.UserResponse{
			User:  nil,
			Error: err.Error(),
		}, nil
	}

	return &gen.UserResponse{
		User: &gen.User{
			Id:    userID,
			Name:  req.Name,
			Email: req.Email,
		},
		Error: "",
	}, nil
}

// GetUser handles fetching a user
func (h *UserHandler) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.UserResponse, error) {
	h.logger.Printf("GetUser called for ID: %s", req.Id)

	user, err := h.service.GetUser(req.Id)
	if err != nil {
		h.logger.Printf("Failed to get user: %v", err)
		return &gen.UserResponse{
			User:  nil,
			Error: err.Error(),
		}, nil
	}

	return &gen.UserResponse{
		User: &gen.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Error: "",
	}, nil
}
