package ports

import (
	"context"

	"github.com/nmdra/Go-gRPC-Sample-Project/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
}
