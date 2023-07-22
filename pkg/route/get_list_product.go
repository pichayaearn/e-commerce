package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	productModel "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type GetProductCfgs struct {
	ProductSvc productModel.ProductSvc
}

func GetListProducts(cfg GetProductCfgs) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.GetListProductsReq)
		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		var productID uuid.UUID
		if req.ProductID != "" {
			productID = uuid.MustParse(req.ProductID)

		}

		listProducts, err := cfg.ProductSvc.ListProduct(productModel.GetProductOpts{
			ProductID: productID,
			Status:    productModel.ProductStatusActive,
		}, c.Request().Context())
		if err != nil {
			return err
		}

		resp := []serializer.GetListProductsResponse{}

		for _, v := range listProducts {
			product := serializer.ToListProductResponse(v)
			resp = append(resp, product)
		}

		return c.JSON(http.StatusOK, resp)

	}
}
