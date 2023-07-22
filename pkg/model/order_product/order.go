package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	OrderStatusActived  OrderStatus = "active"
	OrderStatusCanceled OrderStatus = "canceled"
)

type Order struct {
	id         int
	orderID    uuid.UUID
	userID     uuid.UUID
	total      decimal.Decimal
	status     OrderStatus
	createdAt  time.Time
	updatedAt  time.Time
	canceledAt time.Time
	items      []OrderItem
}

func (o Order) ID() int                { return o.id }
func (o Order) OrderID() uuid.UUID     { return o.orderID }
func (o Order) UserID() uuid.UUID      { return o.userID }
func (o Order) Total() decimal.Decimal { return o.total }
func (o Order) Status() OrderStatus    { return o.status }
func (o Order) CreatedAt() time.Time   { return o.createdAt }
func (o Order) UdatedAt() time.Time    { return o.updatedAt }
func (o Order) CanceledAt() time.Time  { return o.canceledAt }
func (o Order) Items() []OrderItem     { return o.items }

func (o *Order) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&o.orderID, validator.Required),
		validator.Field(&o.status, validator.Required, validator.In(OrderStatusActived, OrderStatusCanceled)),
		validator.Field(&o.createdAt, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(o, rules...); err != nil {
		return err
	}

	return nil
}

func NewOrder() (*Order, error) {
	now := time.Now()
	orderID := uuid.New()
	order := Order{
		orderID:   orderID,
		status:    OrderStatusActived,
		createdAt: now,
		updatedAt: now,
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *Order) SetUserID(userID uuid.UUID) error {
	o.userID = userID
	if err := o.Validate(validator.Field(&o.userID, validator.Required)); err != nil {
		return err
	}
	return nil
}

func (o *Order) SetTotal(total decimal.Decimal) error {
	o.total = total
	if err := o.Validate(validator.Field(&o.total, validator.Required)); err != nil {
		return err
	}
	return nil
}

func (o *Order) SetItems(items []OrderItem) error {
	o.items = items
	if err := o.Validate(); err != nil {
		return err
	}
	return nil
}

func (o *Order) SetCanceled() error {
	o.status = OrderStatusCanceled
	if err := o.Validate(); err != nil {
		return err
	}
	return nil
}
