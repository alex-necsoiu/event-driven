package user

import (
	"context"
	"log"

	"github.com/alex-necsoiu/event-driven/api/proto/gen"
)

// UserHandler implements the gRPC UserServiceServer
type UserHandler struct {
	gen.UnimplementedUserServiceServer
	cfg    Config
	logger *log.Logger
	// Add DB and event bus clients here
}

func NewUserHandler(cfg Config, logger *log.Logger) *UserHandler {
	return &UserHandler{cfg: cfg, logger: logger}
}

// CreateUser handles user creation and publishes an event (placeholder)
func (h *UserHandler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.UserResponse, error) {
	h.logger.Println("CreateUser called (placeholder)")
	// TODO: Insert user into DB, publish UserCreated event
	return &gen.UserResponse{
		User: &gen.User{
			Id:    "1",
			Name:  req.Name,
			Email: req.Email,
		},
		Error: "",
	}, nil
}

// GetUser handles fetching a user (placeholder)
func (h *UserHandler) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.UserResponse, error) {
	h.logger.Println("GetUser called (placeholder)")
	// TODO: Fetch user from DB
	return &gen.UserResponse{
		User: &gen.User{
			Id:    req.Id,
			Name:  "John Doe",
			Email: "john@example.com",
		},
		Error: "",
	}, nil
}
