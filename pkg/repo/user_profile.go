package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	uModel "github.com/pichayaearn/e-commerce/pkg/model/user"
	"github.com/uptrace/bun"
)

type userProfileBun struct {
	bun.BaseModel `bun:"table:user.users_profiles"`
	ID            int `bun:"id,pk,autoincrement"`
	UserID        uuid.UUID
	DisplayName   string
	ProfileImage  string
	Birthday      time.Time
	Gender        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type UserProfileRepo struct {
	db *bun.DB
}

func NewUserProfileRepo(db *bun.DB) uModel.UserProfileRepo {
	return &UserProfileRepo{db: db}
}

func (upr UserProfileRepo) Get(opts uModel.GetUserProfileOpts, ctx context.Context) (*uModel.UserProfile, error) {
	userProfile := userProfileBun{}
	if err := upr.db.NewSelect().Model(&userProfile).ApplyQueryBuilder(addFilter(opts)).OrderExpr("id DESC").Limit(1).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get user error")
	}
	return userProfile.toUserModel()

}

func (upr UserProfileRepo) Create(userProfile uModel.UserProfile) error {
	up := toUserProfileBun(userProfile)
	if _, err := upr.db.NewInsert().Model(&up).Exec(context.Background()); err != nil {
		return errors.New("create user profile failed")
	}
	return nil
}

func addFilter(opts uModel.GetUserProfileOpts) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		if opts.UserID != uuid.Nil {
			q.Where("user_id = ?", opts.UserID)
		}
		return q

	}

}

func (up userProfileBun) toUserModel() (*uModel.UserProfile, error) {
	return uModel.UserProfileFactory(uModel.UserProfileFactoryOpts{
		ID:           up.ID,
		UserID:       up.UserID,
		DisplayName:  up.DisplayName,
		ProfileImage: up.ProfileImage,
		Birthday:     up.Birthday,
		Gender:       up.Gender,
		CreatedAt:    up.CreatedAt,
		UpdatedAt:    up.UpdatedAt,
		DeletedAt:    up.DeletedAt,
	})
}

func toUserProfileBun(userProfile uModel.UserProfile) userProfileBun {
	return userProfileBun{
		ID:           userProfile.ID(),
		UserID:       userProfile.UserID(),
		DisplayName:  userProfile.DisplayName(),
		ProfileImage: userProfile.ProfileName(),
		Birthday:     userProfile.Birthday(),
		Gender:       string(userProfile.Gender()),
		CreatedAt:    userProfile.CreatedAt(),
		UpdatedAt:    userProfile.UpdatedAt(),
		DeletedAt:    userProfile.DeletedAt(),
	}
}
