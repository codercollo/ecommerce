package grpc

import (
	"context"

	"github.com/codercollo/ecommerce/golang/order"
	"github.com/codercollo/ecommerce/order/internal/application/core/domain"
)

// Create handles the gRPC CreateOrder request, translating Protobug types to domain models.
func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem

	// Convert Protobuf order items into domain entities
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	// Instantiate new domain order and execure core application logic
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	// Map domain execution result back to gRPC response payload
	return &order.CreateOrderResponse{
		OrderId: result.ID,
	}, nil
}
