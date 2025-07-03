package order

import (
	"context"
	"log"

	"github.com/alex-necsoiu/event-driven/api/proto/gen"
)

// OrderHandler implements the gRPC OrderServiceServer
type OrderHandler struct {
	gen.UnimplementedOrderServiceServer
	service *Service
	logger  *log.Logger
}

func NewOrderHandler(service *Service, logger *log.Logger) *OrderHandler {
	return &OrderHandler{service: service, logger: logger}
}

// CreateOrder handles order creation and publishes an event
func (h *OrderHandler) CreateOrder(ctx context.Context, req *gen.CreateOrderRequest) (*gen.OrderResponse, error) {
	h.logger.Printf("CreateOrder called for user: %s, amount: %.2f", req.UserId, req.Amount)

	orderID, err := h.service.CreateOrder(req.UserId, req.Amount)
	if err != nil {
		h.logger.Printf("Failed to create order: %v", err)
		return &gen.OrderResponse{
			Order: nil,
			Error: err.Error(),
		}, nil
	}

	return &gen.OrderResponse{
		Order: &gen.Order{
			Id:     orderID,
			UserId: req.UserId,
			Amount: req.Amount,
		},
		Error: "",
	}, nil
}

// GetOrder handles fetching an order
func (h *OrderHandler) GetOrder(ctx context.Context, req *gen.GetOrderRequest) (*gen.OrderResponse, error) {
	h.logger.Printf("GetOrder called for ID: %s", req.Id)

	order, err := h.service.GetOrder(req.Id)
	if err != nil {
		h.logger.Printf("Failed to get order: %v", err)
		return &gen.OrderResponse{
			Order: nil,
			Error: err.Error(),
		}, nil
	}

	return &gen.OrderResponse{
		Order: &gen.Order{
			Id:     order.ID,
			UserId: order.UserID,
			Amount: order.Amount,
		},
		Error: "",
	}, nil
}
