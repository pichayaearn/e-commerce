package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type Gender string

const (
	female Gender = "female"
	male   Gender = "male"
	other  Gender = "other"
)

type UserProfile struct {
	id           int
	userID       uuid.UUID
	displayName  string
	profileImage string
	birthday     time.Time
	gender       Gender
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    time.Time
}

func (up UserProfile) ID() int              { return up.id }
func (up UserProfile) UserID() uuid.UUID    { return up.userID }
func (up UserProfile) DisplayName() string  { return up.displayName }
func (up UserProfile) ProfileName() string  { return up.profileImage }
func (up UserProfile) Birthday() time.Time  { return up.birthday }
func (up UserProfile) Gender() Gender       { return up.gender }
func (up UserProfile) CreatedAt() time.Time { return up.createdAt }
func (up UserProfile) UpdatedAt() time.Time { return up.updatedAt }
func (up UserProfile) DeletedAt() time.Time { return up.deletedAt }

func (up *UserProfile) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&up.userID, validator.Required, is.UUIDv4),
		validator.Field(&up.gender, validator.In(female, male, other)),
		validator.Field(&up.createdAt, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(up, rules...); err != nil {
		return err
	}

	return nil
}

func NewUserProfile(userID uuid.UUID, birthday time.Time, displayName, gender string) (*UserProfile, error) {
	now := time.Now()
	userProfile := UserProfile{
		userID:      userID,
		displayName: displayName,
		birthday:    birthday,
		gender:      Gender(gender),
		createdAt:   now,
		updatedAt:   now,
	}

	if err := userProfile.Validate(); err != nil {
		return nil, err
	}

	return &userProfile, nil
}
