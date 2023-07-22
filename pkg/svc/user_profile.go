package svc

import (
	"context"

	"github.com/google/uuid"
	model "github.com/pichayaearn/e-commerce/pkg/model/user"
)

type UserProfileService struct {
	userProfileRepo model.UserProfileRepo
}

type NewUserProfileSvcCfgs struct {
	UserProfileRepo model.UserProfileRepo
}

func NewUserProfileSvc(cfg NewUserProfileSvcCfgs) model.UserProfileSvc {
	return &UserProfileService{
		userProfileRepo: cfg.UserProfileRepo,
	}
}

func (upSvc UserProfileService) FindProfileByUserID(userID uuid.UUID, ctx context.Context) (*model.UserProfile, error) {
	userProfile, err := upSvc.userProfileRepo.Get(model.GetUserProfileOpts{
		UserID: userID,
	}, ctx)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
