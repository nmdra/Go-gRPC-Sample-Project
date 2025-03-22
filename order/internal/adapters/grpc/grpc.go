package grpc

import (
	"context"

	"github.com/nmdra/Go-gRPC-Sample-Project/order/internal/application/core/domain"
	"github.com/nmdra/Go-gRPC-Sample-Project/order/pb"
	log "github.com/sirupsen/logrus"
)

func (a Adapter) Create(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.WithContext(ctx).Info("Creating order...")
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	result, err := a.api.GetOrder(ctx, request.OrderId)
	var orderItems []*pb.OrderItem
	for _, orderItem := range result.OrderItems {
		orderItems = append(orderItems, &pb.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderResponse{UserId: result.CustomerID, OrderItems: orderItems}, nil
}
