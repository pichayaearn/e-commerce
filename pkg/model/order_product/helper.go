package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductFactoryOpts struct {
	ID        int
	ProductID uuid.UUID
	Name      string
	Amount    int
	Price     decimal.Decimal
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func ProductFactory(opts ProductFactoryOpts) (*Product, error) {
	product := Product{
		id:        opts.ID,
		productID: opts.ProductID,
		name:      opts.Name,
		amount:    opts.Amount,
		price:     opts.Price,
		status:    ProductStatus(opts.Status),
		createdAt: opts.CreatedAt,
		updatedAt: opts.UpdatedAt,
		deletedAt: opts.DeletedAt,
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}

	return &product, nil
}

type OrderFactoryOpts struct {
	ID         int
	OrderID    uuid.UUID
	UserID     uuid.UUID
	Total      decimal.Decimal
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CanceledAt time.Time
}

func OrderFactory(opts OrderFactoryOpts) (*Order, error) {
	order := Order{
		id:         opts.ID,
		orderID:    opts.OrderID,
		userID:     opts.UserID,
		total:      opts.Total,
		status:     OrderStatus(opts.Status),
		createdAt:  opts.CreatedAt,
		updatedAt:  opts.UpdatedAt,
		canceledAt: opts.CanceledAt,
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	return &order, nil
}

type OrderItemFactoryOpts struct {
	ID        int
	OrderID   uuid.UUID
	ProductID uuid.UUID
	Amount    int
	Total     decimal.Decimal
	CreatedAt time.Time
	UpdatedAt time.Time
}

func OrderItemFactory(opts OrderItemFactoryOpts) (*OrderItem, error) {
	orderItem := OrderItem{
		id:        opts.ID,
		orderID:   opts.OrderID,
		productID: opts.ProductID,
		amount:    opts.Amount,
		total:     opts.Total,
		createdAt: opts.CreatedAt,
		updatedAt: opts.UpdatedAt,
	}
	if err := orderItem.Validate(); err != nil {
		return nil, err
	}

	return &orderItem, nil
}
