package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type OrderItemRepo struct {
	db *bun.DB
}

type orderItemsBun struct {
	bun.BaseModel `bun:"table:order.order_items"`
	ID            int `bun:"id,pk,autoincrement"`
	OrderID       uuid.UUID
	ProductID     uuid.UUID
	Amount        int
	Total         decimal.Decimal
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewOrderItemRepo(db *bun.DB) model.OrderItemsRepo {
	return &OrderItemRepo{db: db}
}

func (odir OrderItemRepo) Create(orderItems model.OrderItem) error {
	orderItemBun := toOrderItemBun(orderItems)
	if _, err := odir.db.NewInsert().Model(&orderItemBun).Exec(context.Background()); err != nil {
		return errors.New("create order items failed")
	}
	return nil
}

func (odir OrderItemRepo) List(opts model.GetOrderItems, ctx context.Context) ([]model.OrderItem, error) {
	orderItems := []orderItemsBun{}
	if err := odir.db.NewSelect().Model(&orderItems).ApplyQueryBuilder(addOrderItemsFilter(opts)).OrderExpr("id DESC").Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get list order items error")
	}

	resp := []model.OrderItem{}
	for _, v := range orderItems {
		orderItem, err := v.toOrderItemModel()
		if err != nil {
			return nil, err
		}
		resp = append(resp, *orderItem)
	}

	return resp, nil

}

func addOrderItemsFilter(opts model.GetOrderItems) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		if opts.OrderID != uuid.Nil {
			q.Where("order_id = ?", opts.OrderID)
		}
		return q
	}

}

func toOrderItemBun(orderItems model.OrderItem) orderItemsBun {
	return orderItemsBun{
		ID:        orderItems.ID(),
		OrderID:   orderItems.OrderID(),
		ProductID: orderItems.ProductID(),
		Amount:    orderItems.Amount(),
		Total:     orderItems.Total(),
		CreatedAt: orderItems.CreatedAt(),
		UpdatedAt: orderItems.UpdatedAt(),
	}
}

func (odib orderItemsBun) toOrderItemModel() (*model.OrderItem, error) {
	return model.OrderItemFactory(model.OrderItemFactoryOpts{
		ID:        odib.ID,
		OrderID:   odib.OrderID,
		ProductID: odib.ProductID,
		Amount:    odib.Amount,
		Total:     odib.Total,
		CreatedAt: odib.CreatedAt,
		UpdatedAt: odib.UpdatedAt,
	})
}
