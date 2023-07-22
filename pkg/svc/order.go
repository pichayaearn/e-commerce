package svc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	userModel "github.com/pichayaearn/e-commerce/pkg/model/user"
	"github.com/shopspring/decimal"
)

type OrderSvc struct {
	orderRepo      model.OrderRepo
	orderItemsRepo model.OrderItemsRepo
	productSvc     model.ProductSvc
	userSvc        userModel.UserSvc
}

type NewOrderSvcCfg struct {
	OrderRepo      model.OrderRepo
	OrderItemsRepo model.OrderItemsRepo
	ProductSvc     model.ProductSvc
	UserSvc        userModel.UserSvc
}

func NewOrderSvc(cfg NewOrderSvcCfg) model.OrderSvc {
	return &OrderSvc{
		orderRepo:      cfg.OrderRepo,
		orderItemsRepo: cfg.OrderItemsRepo,
		productSvc:     cfg.ProductSvc,
		userSvc:        cfg.UserSvc,
	}
}

func (odSvc OrderSvc) CreateOrder(opts model.CreateOrder) error {
	ctx := context.Background()
	//find user is active
	userExist, err := odSvc.userSvc.GetUser(userModel.GetUserOpts{
		UserID: opts.UserID,
		Status: userModel.UserStatusActived,
	}, ctx)
	if err != nil {
		return err
	}
	if userExist == nil {
		return fmt.Errorf("user id %s not found", opts.UserID.String())
	}
	// find product is active
	orderItems := []model.OrderItem{}
	listItems := []model.Product{}
	var sumTotalPrice decimal.Decimal
	for _, v := range opts.ProductItems {
		productExist, err := odSvc.productSvc.GetProduct(model.GetProductOpts{
			ProductID: v.ProductID,
			Status:    model.ProductStatusActive,
		}, ctx)
		if err != nil {
			return err
		}

		if productExist == nil || productExist.Amount() <= 0 {
			return fmt.Errorf("product id %s not found", v.ProductID.String())
		}

		if productExist.Amount() < v.Amount {
			return errors.New("not enough amount.")
		}

		//decrease amount
		if err := productExist.DecreaseAmount(v.Amount); err != nil {
			return err
		}

		listItems = append(listItems, *productExist)

		//create new item order
		total := productExist.Price().Mul(decimal.NewFromInt(int64(v.Amount)))
		sumTotalPrice = sumTotalPrice.Add(total)

		newOrderItem := model.NewOrderItems(v.ProductID, v.Amount, total)
		orderItems = append(orderItems, newOrderItem)
	}

	//create new order
	newOrder, err := model.NewOrder()
	if err != nil {
		return err
	}
	if err := newOrder.SetUserID(opts.UserID); err != nil {
		return err
	}
	if err := newOrder.SetTotal(sumTotalPrice); err != nil {
		return err
	}

	if err := odSvc.orderRepo.Create(*newOrder); err != nil {
		return err
	}

	//create order items
	for _, item := range orderItems {
		if err := item.SetOrderID(newOrder.OrderID()); err != nil {
			return err
		}
		if err := odSvc.orderItemsRepo.Create(item); err != nil {
			return err
		}
	}

	//dec amount
	for _, product := range listItems {
		if err := odSvc.productSvc.Update(product); err != nil {
			return err
		}
	}

	return nil
}

func (odSvc OrderSvc) List(opts model.GetOrder, ctx context.Context) ([]model.Order, error) {
	orders, err := odSvc.orderRepo.List(opts, ctx)
	if err != nil {
		return nil, err
	}

	for i, v := range orders {
		//find order items
		orderItems, err := odSvc.orderItemsRepo.List(model.GetOrderItems{
			OrderID: v.OrderID(),
		}, ctx)
		if err != nil {
			return nil, err
		}
		for j, item := range orderItems {
			product, err := odSvc.productSvc.GetProduct(model.GetProductOpts{
				ProductID: item.ProductID(),
			}, ctx)
			if err != nil {
				return nil, err
			}
			if err := item.SetProducts(*product); err != nil {
				return nil, err
			}

			orderItems[j] = item

		}

		if err := v.SetItems(orderItems); err != nil {
			return nil, err
		}

		orders[i] = v

	}
	return orders, nil
}

func (odSvc OrderSvc) CancelOrder(orderID, userID uuid.UUID) error {
	order, err := odSvc.orderRepo.Get(model.GetOrder{
		OrderID: orderID,
		UserID:  userID,
	}, context.Background())
	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	if err := order.SetCanceled(); err != nil {
		return err
	}

	if err := odSvc.orderRepo.Update(*order); err != nil {
		return err
	}

	return nil

}
