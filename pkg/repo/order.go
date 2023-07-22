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

type OrderRepo struct {
	db *bun.DB
}

func NewOrederRepo(db *bun.DB) model.OrderRepo {
	return &OrderRepo{db: db}
}

type orderBun struct {
	bun.BaseModel `bun:"table:order.orders"`
	ID            int `bun:"id,pk,autoincrement"`
	OrderID       uuid.UUID
	UserID        uuid.UUID
	Total         decimal.Decimal
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CanceledAt    time.Time
}

func (or OrderRepo) Create(order model.Order) error {
	orderBun := toOrderBun(order)
	if _, err := or.db.NewInsert().Model(&orderBun).Exec(context.Background()); err != nil {
		return errors.New("create order failed")
	}
	return nil
}

func (or OrderRepo) Get(opts model.GetOrder, ctx context.Context) (*model.Order, error) {
	order := orderBun{}
	if err := or.db.NewSelect().Model(&order).ApplyQueryBuilder(addOrderFilter(opts)).OrderExpr("id DESC").Limit(1).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get list order error")
	}

	return order.toOrderModel()

}

func (or OrderRepo) List(opts model.GetOrder, ctx context.Context) ([]model.Order, error) {
	orders := []orderBun{}
	if err := or.db.NewSelect().Model(&orders).ApplyQueryBuilder(addOrderFilter(opts)).OrderExpr("id DESC").Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get list order error")
	}

	resp := []model.Order{}

	for _, v := range orders {
		order, err := v.toOrderModel()
		if err != nil {
			return nil, err
		}
		resp = append(resp, *order)
	}

	return resp, nil

}

func (or OrderRepo) Update(order model.Order) error {
	pd := toOrderBun(order)
	if _, err := or.db.NewUpdate().
		Model(&pd).
		WherePK().
		Exec(context.Background()); err != nil {
		return err
	}
	return nil
}

func addOrderFilter(opts model.GetOrder) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		if opts.OrderID != uuid.Nil {
			q.Where("order_id = ?", opts.OrderID)
		}
		if opts.UserID != uuid.Nil {
			q.Where("user_id = ?", opts.UserID)
		}
		if opts.Status != "" {
			q.Where("status = ?", opts.Status)
		}
		return q

	}

}

func toOrderBun(order model.Order) orderBun {
	return orderBun{
		ID:         order.ID(),
		OrderID:    order.OrderID(),
		UserID:     order.UserID(),
		Total:      order.Total(),
		Status:     string(order.Status()),
		CreatedAt:  order.CreatedAt(),
		UpdatedAt:  order.UdatedAt(),
		CanceledAt: order.CanceledAt(),
	}
}

func (ob orderBun) toOrderModel() (*model.Order, error) {
	return model.OrderFactory(model.OrderFactoryOpts{
		ID:         ob.ID,
		OrderID:    ob.OrderID,
		UserID:     ob.UserID,
		Total:      ob.Total,
		Status:     ob.Status,
		CreatedAt:  ob.CreatedAt,
		UpdatedAt:  ob.UpdatedAt,
		CanceledAt: ob.CanceledAt,
	})
}
