package model

import (
	"context"

	"github.com/google/uuid"
)

type GetProductOpts struct {
	ID        int
	ProductID uuid.UUID
	Status    ProductStatus
}

type ProductRepo interface {
	List(opts GetProductOpts, ctx context.Context) ([]Product, error)
	Get(opts GetProductOpts, ctx context.Context) (*Product, error)
	Update(product Product) error
}

type OrderRepo interface {
	Create(order Order) error
	List(opts GetOrder, ctx context.Context) ([]Order, error)
	Get(opts GetOrder, ctx context.Context) (*Order, error)
	Update(order Order) error
}

type GetOrderItems struct {
	OrderID uuid.UUID
}

type OrderItemsRepo interface {
	Create(orderItems OrderItem) error
	List(opts GetOrderItems, ctx context.Context) ([]OrderItem, error)
}
