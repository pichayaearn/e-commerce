package model

import (
	"time"

	"github.com/google/uuid"
)

type UserFactoryOpts struct {
	ID        int
	UserID    uuid.UUID
	Email     string
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func UserFactory(opts UserFactoryOpts) (*User, error) {
	user := User{
		id:        opts.ID,
		userID:    opts.UserID,
		email:     opts.Email,
		password:  opts.Password,
		status:    UserStatus(opts.Status),
		createdAt: opts.CreatedAt,
		updatedAt: opts.UpdatedAt,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil
}

type UserProfileFactoryOpts struct {
	ID           int
	UserID       uuid.UUID
	DisplayName  string
	ProfileImage string
	Birthday     time.Time
	Gender       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func UserProfileFactory(opts UserProfileFactoryOpts) (*UserProfile, error) {
	userProfile := UserProfile{
		id:           opts.ID,
		userID:       opts.UserID,
		displayName:  opts.DisplayName,
		profileImage: opts.ProfileImage,
		birthday:     opts.Birthday,
		gender:       Gender(opts.Gender),
		createdAt:    opts.CreatedAt,
		updatedAt:    opts.UpdatedAt,
		deletedAt:    opts.DeletedAt,
	}

	if err := userProfile.Validate(); err != nil {
		return nil, err
	}

	return &userProfile, nil
}
