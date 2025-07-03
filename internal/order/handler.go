package order

import (
	"context"
	"log"

	"github.com/alex-necsoiu/event-driven/api/proto/gen"
)

// OrderHandler implements the gRPC OrderServiceServer
type OrderHandler struct {
	cfg    Config
	logger *log.Logger
	// Add DB and event bus clients here
}

func NewOrderHandler(cfg Config, logger *log.Logger) *OrderHandler {
	return &OrderHandler{cfg: cfg, logger: logger}
}

// CreateOrder handles order creation and publishes an event (placeholder)
func (h *OrderHandler) CreateOrder(ctx context.Context, req *gen.CreateOrderRequest) (*gen.OrderResponse, error) {
	h.logger.Println("CreateOrder called (placeholder)")
	// TODO: Insert order into DB, publish OrderCreated event
	return &gen.OrderResponse{
		Order: &gen.Order{
			Id:     "1",
			UserId: req.UserId,
			Amount: req.Amount,
		},
		Error: "",
	}, nil
}

// GetOrder handles fetching an order (placeholder)
func (h *OrderHandler) GetOrder(ctx context.Context, req *gen.GetOrderRequest) (*gen.OrderResponse, error) {
	h.logger.Println("GetOrder called (placeholder)")
	// TODO: Fetch order from DB
	return &gen.OrderResponse{
		Order: &gen.Order{
			Id:     req.Id,
			UserId: "1",
			Amount: 99.99,
		},
		Error: "",
	}, nil
}

// mustEmbedUnimplementedOrderServiceServer implements the gRPC interface requirement
func (h *OrderHandler) mustEmbedUnimplementedOrderServiceServer() {}
