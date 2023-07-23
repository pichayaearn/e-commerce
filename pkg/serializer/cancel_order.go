package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type CancelOrderReq struct {
	OrderID string `json:"order_id"`
	UserID  uuid.UUID
}

func NewCancelOrderReq(userID uuid.UUID) *CancelOrderReq {
	return &CancelOrderReq{
		UserID: userID,
	}
}

func (r CancelOrderReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required, is.UUIDv4),
		validation.Field(&r.OrderID, validation.Required, is.UUIDv4),
	)
}
