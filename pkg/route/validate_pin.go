package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

func ValidatePin() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.ValidatePinReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}
