package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	// jwt.RegisteredClaims
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.LoginReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		// Create a JWT token
		// claims := &jwtCustomClaims{
		// 	Email: req.Email,
		// 	jwt.RegisteredClaims{
		// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		// 	},
		// }

		return c.NoContent(200)

	}
}
