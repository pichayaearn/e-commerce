package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type CreateOrderCfgs struct {
	OrderSvc model.OrderSvc
}

func CreateOrder(cfg CreateOrderCfgs) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewCreateOrderReq(userID)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		items := []model.ProductItems{}
		for _, v := range req.Items {
			productID := uuid.MustParse(v.ProductID)
			items = append(items, model.ProductItems{
				ProductID: productID,
				Amount:    v.Amount,
			})
		}

		if err := cfg.OrderSvc.CreateOrder(model.CreateOrder{
			UserID:       userID,
			ProductItems: items,
		}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Create Order: "+err.Error())
		}

		return c.NoContent(http.StatusCreated)

	}
}
