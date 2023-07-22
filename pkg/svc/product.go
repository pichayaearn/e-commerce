package svc

import (
	"context"

	productModel "github.com/pichayaearn/e-commerce/pkg/model/order_product"
)

type ProductSvc struct {
	productRepo productModel.ProductRepo
}

type NewProductSvcCfgs struct {
	ProductRepo productModel.ProductRepo
}

func NewProductSvc(cfg NewProductSvcCfgs) productModel.ProductSvc {
	return &ProductSvc{productRepo: cfg.ProductRepo}
}

func (pdSvc ProductSvc) ListProduct(opts productModel.GetProductOpts, ctx context.Context) ([]productModel.Product, error) {
	return pdSvc.productRepo.List(opts, ctx)
}

func (pdSvc ProductSvc) GetProduct(opts productModel.GetProductOpts, ctx context.Context) (*productModel.Product, error) {
	return pdSvc.productRepo.Get(opts, ctx)
}

func (pdSvc ProductSvc) Update(product productModel.Product) error {
	if err := pdSvc.productRepo.Update(product); err != nil {
		return err
	}

	return nil
}
