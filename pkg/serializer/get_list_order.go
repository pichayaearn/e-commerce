package serializer

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/shopspring/decimal"
)

type GetListOrderReq struct {
	OrderID string `json:"order_id" query:"order_id"`

	UserID uuid.UUID
}

func (r GetListOrderReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required, is.UUIDv4),
		validation.Field(&r.OrderID, is.UUIDv4),
	)
}

func NewGetListOrderReq(userID uuid.UUID) *GetListOrderReq {
	return &GetListOrderReq{
		UserID: userID,
	}
}

type GetListOrderResponse struct {
	OrderID   uuid.UUID
	Total     decimal.Decimal
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []ListItem
}

type ListItem struct {
	ProductID uuid.UUID       `json:"product_id"`
	Amount    int             `json:"amount"`
	Name      string          `json:"name"`
	Price     decimal.Decimal `json:"price"`
}

func ToListOrderResponse(order model.Order) GetListOrderResponse {
	resp := GetListOrderResponse{
		OrderID:   order.OrderID(),
		Total:     order.Total(),
		Status:    string(order.Status()),
		CreatedAt: order.CreatedAt(),
		UpdatedAt: order.UdatedAt(),
	}

	items := []ListItem{}
	for _, v := range order.Items() {
		item := ListItem{
			ProductID: v.ProductID(),
			Amount:    v.Amount(),
			Name:      v.Product().Name(),
			Price:     v.Product().Price(),
		}
		items = append(items, item)
	}

	resp.Items = items

	return resp
}
