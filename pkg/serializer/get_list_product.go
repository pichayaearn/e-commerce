package serializer

import (
	"time"

	"github.com/google/uuid"
	productModel "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/shopspring/decimal"
)

type GetListProductsReq struct {
	ProductID string `json:"id" query:"product_id"`
}

type GetListProductsResponse struct {
	ID        uuid.UUID       `json:"product_id"`
	Name      string          `json:"name"`
	Amount    int             `json:"amount"`
	Price     decimal.Decimal `json:"price"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func ToListProductResponse(product productModel.Product) GetListProductsResponse {
	return GetListProductsResponse{
		ID:        product.ProductID(),
		Name:      product.Name(),
		Amount:    product.Amount(),
		Price:     product.Price(),
		Status:    string(product.Status()),
		CreatedAt: product.CreatedAt(),
		UpdatedAt: product.UdatedAt(),
	}
}
