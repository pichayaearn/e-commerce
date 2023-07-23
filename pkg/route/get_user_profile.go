package route

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	model "github.com/pichayaearn/e-commerce/pkg/model/user"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
)

type GetUserProfileCfg struct {
	UserProfileSvc model.UserProfileSvc
}

func GetUserProfile(cfg GetUserProfileCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewUserProfileRequest(userID)

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

func BindUserIDFromContext(c echo.Context) (uuid.UUID, error) {
	userID := c.Get("UserID")
	userIDStr, ok := userID.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user id not string")
	}

	uid := uuid.MustParse(userIDStr)
	return uid, nil
}
