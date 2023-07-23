package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type CreateOrderReq struct {
	Items  []Item `json:"items"`
	UserID uuid.UUID
}
type Item struct {
	ProductID string `json:"product_id"`
	Amount    int    `json:"amount"`
}

func NewCreateOrderReq(userID uuid.UUID) *CreateOrderReq {
	return &CreateOrderReq{
		UserID: userID,
	}
}

func (r CreateOrderReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required),
	)
}
