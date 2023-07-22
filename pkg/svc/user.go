package svc

import (
	"context"
	"errors"

	model "github.com/pichayaearn/e-commerce/pkg/model/user"
)

type UserSvc struct {
	userRepo        model.UserRepo
	userProfileRepo model.UserProfileRepo
}

type NewUserSvcCfgs struct {
	UserRepo        model.UserRepo
	UserProfileRepo model.UserProfileRepo
}

func NewUserSvc(cfg NewUserSvcCfgs) model.UserSvc {
	return &UserSvc{
		userRepo:        cfg.UserRepo,
		userProfileRepo: cfg.UserProfileRepo,
	}
}

func (uSvc UserSvc) CreateUser(opts model.CreateUser) (*model.User, *model.UserProfile, error) {
	//check email exist
	userExist, err := uSvc.userRepo.Get(model.GetUserOpts{
		Email: opts.Email,
	}, context.Background())
	if err != nil {
		return nil, nil, err
	}

	if userExist != nil {
		//email already used
		return nil, nil, errors.New("email is already used")
	}

	//create user
	newUser, err := model.NewUser(opts.Email, opts.Password)
	if err != nil {
		return nil, nil, err
	}
	if err := uSvc.userRepo.Create(*newUser); err != nil {
		return nil, nil, err
	}

	//create user profile
	newUserProfile, err := model.NewUserProfile(newUser.UserID(), opts.Birthday, opts.DisplayName, opts.Gender)
	if err != nil {
		return nil, nil, err
	}
	if err := uSvc.userProfileRepo.Create(*newUserProfile); err != nil {
		return nil, nil, err
	}

	return newUser, newUserProfile, nil
}

func (uSvc UserSvc) GetUser(opts model.GetUserOpts, ctx context.Context) (*model.User, error) {
	return uSvc.userRepo.Get(opts, ctx)
}
