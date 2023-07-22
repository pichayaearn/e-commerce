package serializer

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	model "github.com/pichayaearn/e-commerce/pkg/model/user"
)

type CreateUserReq struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DisplayName string    `json:"display_name"`
	Birthday    string    `json:"birthday"`
	Gender      string    `json:"gender"`
	Birthdate   time.Time `json:"-"`
}

func (r *CreateUserReq) ParseBirthdate() error {
	//parse birthdate to type time
	birthdate, err := time.Parse(time.RFC3339, r.Birthday+"T00:00:00Z")
	if err != nil {
		return err
	}
	r.Birthdate = birthdate
	return nil
}

func (r CreateUserReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
		validation.Field(&r.DisplayName, validation.Required),
		validation.Field(&r.Birthdate, validation.Required),
		validation.Field(&r.Gender, validation.Required),
	)
}

type CreateUserResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	Birthday    time.Time `json:"birthday"`
	Gender      string    `json:"gender"`
}

func ToCreateUserResponse(user model.User, userProfile model.UserProfile) CreateUserResponse {
	return CreateUserResponse{
		UserID:      user.UserID(),
		Email:       user.Email(),
		DisplayName: userProfile.DisplayName(),
		Birthday:    userProfile.Birthday(),
		Gender:      string(userProfile.Gender()),
	}
}
