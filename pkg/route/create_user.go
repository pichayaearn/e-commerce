package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/pichayaearn/e-commerce/pkg/model/user"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type CreateUserCfg struct {
	UserSvc model.UserSvc
}

func CreateUser(cfg CreateUserCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.CreateUserReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//parsed birthdate
		if err := req.ParseBirthdate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		user, userProfile, err := cfg.UserSvc.CreateUser(model.CreateUser{
			Email:       req.Email,
			Password:    req.Password,
			DisplayName: req.DisplayName,
			Birthday:    req.Birthdate,
			Gender:      req.Gender,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Create user"+err.Error())
		}

		resp := serializer.ToCreateUserResponse(*user, *userProfile)

		return c.JSON(http.StatusOK, resp)

	}
}
