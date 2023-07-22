package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateOrderReq struct {
	UserID string `json:"user_id"`
	Items  []Item `json:"items"`
}
type Item struct {
	ProductID string `json:"product_id"`
	Amount    int    `json:"amount"`
}

func (r CreateOrderReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required),
	)
}
