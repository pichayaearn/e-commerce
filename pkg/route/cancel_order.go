package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type CancelOrderCfgs struct {
	OrderSvc model.OrderSvc
}

func CancelOrder(cfg CancelOrderCfgs) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewCancelOrderReq(userID)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		var orderID uuid.UUID
		if req.OrderID != "" {
			orderID = uuid.MustParse(req.OrderID)

		}

		if err := cfg.OrderSvc.CancelOrder(orderID, userID); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Canceled Order: "+err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}
