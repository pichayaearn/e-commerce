package route

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	model "github.com/pichayaearn/e-commerce/pkg/model/order_product"
	"github.com/pichayaearn/e-commerce/pkg/model/order_product/mocks"
	"github.com/pichayaearn/e-commerce/pkg/serializer"
	"github.com/stretchr/testify/require"
)

func TestGetListorder(t *testing.T) {
	r := require.New(t)

	type want struct {
		isError bool
	}

	testCases := []struct {
		name    string
		request serializer.GetListOrderReq
		mockSvc func(m *mocks.OrderSvc)
		want    want
	}{
		{
			name: "failed, user id is nil",
			request: serializer.GetListOrderReq{
				UserID: uuid.Nil,
			},
			want: want{
				isError: true,
			},
		},
		{
			name: "failed, validate req failed",
			request: serializer.GetListOrderReq{
				UserID:  uuid.MustParse("a2dbd5f5-fc78-48cc-b708-a329747f6b33"),
				OrderID: "test",
			},
			want: want{
				isError: true,
			},
		},
		{
			name: "failed, find order list",
			request: serializer.GetListOrderReq{
				UserID: uuid.MustParse("a2dbd5f5-fc78-48cc-b708-a329747f6b33"),
			},
			mockSvc: func(m *mocks.OrderSvc) {
				userID := uuid.MustParse("a2dbd5f5-fc78-48cc-b708-a329747f6b33")
				m.On("List", model.GetOrder{
					UserID: userID,
				}, context.Background()).Return(nil, errors.New("error"))
			},
			want: want{
				isError: true,
			},
		},
		{
			name: "success",
			request: serializer.GetListOrderReq{
				UserID: uuid.MustParse("a2dbd5f5-fc78-48cc-b708-a329747f6b33"),
			},
			mockSvc: func(m *mocks.OrderSvc) {
				userID := uuid.MustParse("a2dbd5f5-fc78-48cc-b708-a329747f6b33")
				order, err := model.NewOrder()
				r.NoError(err)
				list := []model.Order{}
				list = append(list, *order)
				m.On("List", model.GetOrder{
					UserID: userID,
				}, context.Background()).Return(list, nil)
			},
			want: want{
				isError: false,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.request)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodGet, "/orders", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)

			if tt.request.UserID != uuid.Nil {
				c.Set("UserID", tt.request.UserID.String())
			}

			mockSvc := mocks.NewOrderSvc(t)

			if tt.mockSvc != nil {
				tt.mockSvc(mockSvc)
			}

			err = GetListorder(GetListOrderCfgs{
				OrderSvc: mockSvc,
			})(c)

			if tt.want.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.Equal(http.StatusOK, rec.Code)
			}

		})
	}
}
