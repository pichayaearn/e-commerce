package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductStatus string

const (
	ProductStatusActive  ProductStatus = "active"
	ProductStatusDeleted ProductStatus = "deleted"
)

type Product struct {
	id        int
	productID uuid.UUID
	name      string
	amount    int
	price     decimal.Decimal
	status    ProductStatus
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func (p Product) ID() int                { return p.id }
func (p Product) ProductID() uuid.UUID   { return p.productID }
func (p Product) Name() string           { return p.name }
func (p Product) Amount() int            { return p.amount }
func (p Product) Price() decimal.Decimal { return p.price }
func (p Product) Status() ProductStatus  { return p.status }
func (p Product) CreatedAt() time.Time   { return p.createdAt }
func (p Product) UdatedAt() time.Time    { return p.updatedAt }
func (p Product) DeletedAt() time.Time   { return p.deletedAt }

func (p *Product) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&p.id, validator.Required),
		validator.Field(&p.name, validator.Required),
		validator.Field(&p.price, validator.Required),
		validator.Field(&p.status, validator.Required, validator.In(ProductStatusActive, ProductStatusDeleted)),
		validator.Field(&p.createdAt, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(p, rules...); err != nil {
		return err
	}

	return nil
}

func (p *Product) DecreaseAmount(amount int) error {
	p.amount -= amount
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}
