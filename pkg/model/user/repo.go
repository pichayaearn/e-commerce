package model

import (
	"context"

	"github.com/google/uuid"
)

type GetUserOpts struct {
	UserID uuid.UUID
	Email  string
	Status UserStatus
}

type GetUserProfileOpts struct {
	UserID uuid.UUID
}

type UserRepo interface {
	Get(opts GetUserOpts, ctx context.Context) (*User, error)
	Create(user User) error
}

type UserProfileRepo interface {
	Get(opts GetUserProfileOpts, ctx context.Context) (*UserProfile, error)
	Create(userProfile UserProfile) error
}
