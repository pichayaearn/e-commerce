package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateUser struct {
	Email       string
	Password    string
	DisplayName string
	Birthday    time.Time
	Gender      string
}

type UserSvc interface {
	CreateUser(opts CreateUser) (*User, *UserProfile, error)
	GetUser(opts GetUserOpts, ctx context.Context) (*User, error)
}

type UserProfileSvc interface {
	FindProfileByUserID(userID uuid.UUID, ctx context.Context) (*UserProfile, error)
}
