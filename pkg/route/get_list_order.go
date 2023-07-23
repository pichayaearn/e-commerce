package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type GetListOrderCfgs struct {
	OrderSvc model.OrderSvc
}

func GetListorder(cfg GetListOrderCfgs) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewGetListOrderReq(userID)

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

		orders, err := cfg.OrderSvc.List(model.GetOrder{
			OrderID: orderID,
			UserID:  userID,
		}, c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Get list order: "+err.Error())
		}

		resp := []serializer.GetListOrderResponse{}

		for _, v := range orders {
			order := serializer.ToListOrderResponse(v)
			resp = append(resp, order)
		}

		return c.JSON(http.StatusOK, resp)

	}
}
