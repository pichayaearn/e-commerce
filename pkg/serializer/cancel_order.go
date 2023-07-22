package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CancelOrderReq struct {
	UserID  string `json:"user_id"`
	OrderID string `json:"order_id"`
}

func (r CancelOrderReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required, is.UUIDv4),
		validation.Field(&r.OrderID, validation.Required, is.UUIDv4),
	)
}
