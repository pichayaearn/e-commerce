package model

import (
	"context"

	"github.com/google/uuid"
)

type ProductSvc interface {
	ListProduct(opts GetProductOpts, ctx context.Context) ([]Product, error)
	GetProduct(opts GetProductOpts, ctx context.Context) (*Product, error)
	Update(product Product) error
}

type ProductItems struct {
	ProductID uuid.UUID
	Amount    int
}

type CreateOrder struct {
	UserID       uuid.UUID
	ProductItems []ProductItems
}

type GetOrder struct {
	OrderID uuid.UUID
	Status  OrderStatus
	UserID  uuid.UUID
}

type OrderSvc interface {
	CreateOrder(opts CreateOrder) error
	List(opts GetOrder, ctx context.Context) ([]Order, error)
	CancelOrder(orderID, userID uuid.UUID) error
}
