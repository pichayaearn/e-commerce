package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	productModel "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type ProductRepo struct {
	db *bun.DB
}

func NewProductRepo(db *bun.DB) productModel.ProductRepo {
	return &ProductRepo{db: db}
}

type productBun struct {
	bun.BaseModel `bun:"table:product.products"`
	ID            int `bun:"id,pk,autoincrement"`
	ProductID     uuid.UUID
	Name          string
	Amount        int
	Price         decimal.Decimal
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (pr ProductRepo) Get(opts productModel.GetProductOpts, ctx context.Context) (*productModel.Product, error) {
	product := productBun{}
	if err := pr.db.NewSelect().Model(&product).ApplyQueryBuilder(addProductFilter(opts)).OrderExpr("id DESC").Limit(1).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get product error")
	}

	return product.toProductModel()

}

func (pr ProductRepo) List(opts productModel.GetProductOpts, ctx context.Context) ([]productModel.Product, error) {
	products := []productBun{}
	if err := pr.db.NewSelect().Model(&products).ApplyQueryBuilder(addProductFilter(opts)).OrderExpr("name DESC").Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get list product error")
	}

	resp := []productModel.Product{}

	for _, v := range products {
		product, err := v.toProductModel()
		if err != nil {
			return nil, err
		}
		resp = append(resp, *product)
	}

	return resp, nil

}

func (pr ProductRepo) Update(product model.Product) error {
	pd := toProductBun(product)
	if _, err := pr.db.NewUpdate().
		Model(&pd).
		WherePK().
		Exec(context.Background()); err != nil {
		return err
	}
	return nil
}

func addProductFilter(opts productModel.GetProductOpts) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		if opts.ID != 0 {
			q.Where("id = ?", opts.ID)
		}
		if opts.ProductID != uuid.Nil {
			q.Where("product_id = ?", opts.ProductID)
		}
		if opts.Status != "" {
			q.Where("status = ?", opts.Status)
		}
		return q

	}

}

func toProductBun(product model.Product) productBun {
	return productBun{
		ID:        product.ID(),
		ProductID: product.ProductID(),
		Name:      product.Name(),
		Amount:    product.Amount(),
		Price:     product.Price(),
		Status:    string(product.Status()),
		CreatedAt: product.CreatedAt(),
		UpdatedAt: product.UdatedAt(),
		DeletedAt: product.DeletedAt(),
	}
}

func (pb productBun) toProductModel() (*productModel.Product, error) {
	return productModel.ProductFactory(productModel.ProductFactoryOpts{
		ID:        pb.ID,
		ProductID: pb.ProductID,
		Name:      pb.Name,
		Amount:    pb.Amount,
		Price:     pb.Price,
		Status:    pb.Status,
		CreatedAt: pb.CreatedAt,
		UpdatedAt: pb.UpdatedAt,
		DeletedAt: pb.DeletedAt,
	})
}
