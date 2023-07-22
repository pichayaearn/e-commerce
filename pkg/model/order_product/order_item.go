package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderItem struct {
	id        int
	orderID   uuid.UUID
	productID uuid.UUID
	amount    int
	total     decimal.Decimal
	createdAt time.Time
	updatedAt time.Time
	product   Product
}

func (odi OrderItem) ID() int                { return odi.id }
func (odi OrderItem) OrderID() uuid.UUID     { return odi.orderID }
func (odi OrderItem) ProductID() uuid.UUID   { return odi.productID }
func (odi OrderItem) Amount() int            { return odi.amount }
func (odi OrderItem) Total() decimal.Decimal { return odi.total }
func (odi OrderItem) CreatedAt() time.Time   { return odi.createdAt }
func (odi OrderItem) UpdatedAt() time.Time   { return odi.updatedAt }
func (odi OrderItem) Product() Product       { return odi.product }

func (odi *OrderItem) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&odi.productID, validator.Required),
		validator.Field(&odi.createdAt, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(odi, rules...); err != nil {
		return err
	}

	return nil
}

func NewOrderItems(productID uuid.UUID, amount int, total decimal.Decimal) OrderItem {
	now := time.Now()
	return OrderItem{
		productID: productID,
		amount:    amount,
		total:     total,
		createdAt: now,
		updatedAt: now,
	}
}

func (odi *OrderItem) SetOrderID(orderID uuid.UUID) error {
	odi.orderID = orderID
	if err := odi.Validate(validator.Field(&odi.orderID, validator.Required)); err != nil {
		return err
	}
	return nil
}

func (odi *OrderItem) SetProducts(product Product) error {
	odi.product = product
	if err := odi.Validate(); err != nil {
		return err
	}
	return nil
}
