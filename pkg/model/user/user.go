package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusActived UserStatus = "active"
	UserStatusDeleted UserStatus = "deleted"
)

type User struct {
	id        int
	userID    uuid.UUID
	email     string
	password  string
	status    UserStatus
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func (u User) ID() int              { return u.id }
func (u User) UserID() uuid.UUID    { return u.userID }
func (u User) Email() string        { return u.email }
func (u User) Password() string     { return u.password }
func (u User) Status() UserStatus   { return u.status }
func (u User) CreatedAt() time.Time { return u.createdAt }
func (u User) UpdatedAt() time.Time { return u.updatedAt }
func (u User) DeletedAt() time.Time { return u.deletedAt }

func (u *User) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&u.userID, validator.Required, is.UUIDv4),
		validator.Field(&u.email, validator.Required),
		validator.Field(&u.password, validator.Required),
		validator.Field(&u.status, validator.Required, validator.In(UserStatusActived, UserStatusDeleted)),
		validator.Field(&u.createdAt, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(u, rules...); err != nil {
		return err
	}

	return nil
}

func NewUser(email, password string) (*User, error) {
	userID := uuid.New()
	now := time.Now()
	user := User{
		userID:    userID,
		email:     email,
		password:  password,
		status:    UserStatusActived,
		createdAt: now,
		updatedAt: now,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil
}
