package serializer

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	model "github.com/pichayaearn/e-commerce/pkg/model/user"
)

type UserProfileReq struct {
	UserID uuid.UUID `json:"user_id"`
}

func (up UserProfileReq) ValidateRequest() error {
	return validation.ValidateStruct(&up,
		validation.Field(&up.UserID, validation.Required, is.UUIDv4),
	)
}

type UserProfileResponse struct {
	UserID       uuid.UUID `json:"user_id"`
	DisplayName  string    `json:"display_name"`
	ProfileImage string    `json:"profile_image"`
	Birthday     time.Time `json:"birthday"`
	Gender       string    `json:"gender"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToUserProfileResponse(up model.UserProfile) UserProfileResponse {
	return UserProfileResponse{
		UserID:       up.UserID(),
		DisplayName:  up.DisplayName(),
		ProfileImage: up.ProfileName(),
		Birthday:     up.Birthday(),
		Gender:       string(up.Gender()),
		CreatedAt:    up.CreatedAt(),
		UpdatedAt:    up.UpdatedAt(),
	}
}
