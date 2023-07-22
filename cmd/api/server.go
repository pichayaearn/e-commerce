package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/e-commerce/cmd/api/config"
	"github.com/pichayaearn/e-commerce/pkg/repo"
	"github.com/pichayaearn/e-commerce/pkg/route"
	"github.com/pichayaearn/e-commerce/pkg/svc"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun/extra/bundebug"
)

func newServer(cfg *config.Config) *echo.Echo {
	db := cfg.DB.MustNewDB()
	logger := logrus.New()
	logger.Info("new server")
	if cfg.Environment == "development" {
		db.AddQueryHook(bundebug.NewQueryHook())
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	e := echo.New()

	userRepo := repo.NewUserRepo(db)
	userProfileRepo := repo.NewUserProfileRepo(db)
	productRepo := repo.NewProductRepo(db)
	orderRepo := repo.NewOrederRepo(db)
	orderItemRepo := repo.NewOrderItemRepo(db)

	userSvc := svc.NewUserSvc(svc.NewUserSvcCfgs{
		UserRepo:        userRepo,
		UserProfileRepo: userProfileRepo,
	})
	userProfileSvc := svc.NewUserProfileSvc(svc.NewUserProfileSvcCfgs{
		UserProfileRepo: userProfileRepo,
	})
	productSvc := svc.NewProductSvc(svc.NewProductSvcCfgs{
		ProductRepo: productRepo,
	})
	orderSvc := svc.NewOrderSvc(svc.NewOrderSvcCfg{
		OrderRepo:      orderRepo,
		OrderItemsRepo: orderItemRepo,
		ProductSvc:     productSvc,
		UserSvc:        userSvc,
	})

	e.POST("/sign-up", route.CreateUser(route.CreateUserCfg{
		UserSvc: userSvc,
	}))

	e.GET("/user-profile", route.GetUserProfile(route.GetUserProfileCfg{
		UserProfileSvc: userProfileSvc,
	}))

	e.GET("/list-products", route.GetListProducts(route.GetProductCfgs{
		ProductSvc: productSvc,
	}))

	e.POST("/orders", route.CreateOrder(route.CreateOrderCfgs{
		OrderSvc: orderSvc,
	}))

	e.GET("/orders", route.GetListorder(route.GetListOrderCfgs{
		OrderSvc: orderSvc,
	}))

	e.PATCH("/cancel", route.CancelOrder(route.CancelOrderCfgs{
		OrderSvc: orderSvc,
	}))

	return e

}
