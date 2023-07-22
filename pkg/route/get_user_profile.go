package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/pichayaearn/e-commerce/pkg/model/user"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type GetUserProfileCfg struct {
	UserProfileSvc model.UserProfileSvc
}

func GetUserProfile(cfg GetUserProfileCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.UserProfileReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		//get user profile
		userProfile, err := cfg.UserProfileSvc.FindProfileByUserID(req.UserID, c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Get user profile "+err.Error())
		}

		resp := serializer.ToUserProfileResponse(*userProfile)

		return c.JSON(http.StatusOK, resp)

	}
}
