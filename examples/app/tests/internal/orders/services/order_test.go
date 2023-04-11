//go:build unit_test

// Приклад написання юніт тестів для бізнес логіки.
// Такі тести мають виключати будь-які зовнішні підключення і тестувати логіку в ізоляції.
// Важливим інструментом для написання тестів є інтерфейси,
// які дозволяють легко описувати мок дані, як у `order_mock_test.go`.
// Юніт тести також доречно використовувати для хендлерів.

package services

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"tests/internal/orders"
)

func TestGet(t *testing.T) {
	orderGetError := errors.New("failed to get order")
	order := orders.Order{
		UUID:        uuid.New(),
		Date:        time.Now(),
		Description: "test",
		Price:       10,
		Status:      orders.OrderStatus(2),
	}

	// В go рекомендується застосовувати підхід table testing`.
	// Він полягає в тому, що ситуації, що перевіряються описуються
	// як колекція вхідних та вихідних даних.
	// Потім в циклі відбувається прохід по цим даним і відбувається
	// виклик методів та перевірка, однакова для всіх варіантів.
	// Це дозволяє легко редагувати та додавати варіанти перевірок.
	testCases := []struct {
		description     string
		orderGetSuccess *orders.Order
		orderGetError   error
		expectedSuccess *orders.Order
		expectedError   error
	}{
		{
			description:   "get error from order repository",
			orderGetError: orderGetError,
			expectedError: orderGetError,
		},
		{
			description:   "not found error from order repository",
			orderGetError: pgx.ErrNoRows,
			expectedError: ErrNotFound,
		},
		{
			description:     "success order, empty order product list",
			orderGetSuccess: &order,
			expectedSuccess: &order,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			orderRepositoryMock := &OrderRepositoryMock{
				GetSuccess: tc.orderGetSuccess,
				GetError:   tc.orderGetError,
			}

			o := NewOrderService(orderRepositoryMock)
			result, err := o.OrderGet(context.Background(), order.UUID)
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}

			if !reflect.DeepEqual(tc.expectedSuccess, result) {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedSuccess, result)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	orderCreateError := errors.New("failed to create order")
	testCases := []struct {
		description      string
		orderCreateError error
		expectedError    error
	}{
		{
			description:      "get error from order repository",
			orderCreateError: orderCreateError,
			expectedError:    orderCreateError,
		},
		{
			description: "success order create",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			orderRepositoryMock := &OrderRepositoryMock{
				CreateError: tc.orderCreateError,
			}

			o := NewOrderService(orderRepositoryMock)
			err := o.OrderCreate(context.Background(), &orders.Order{
				UUID:        uuid.New(),
				Date:        time.Now(),
				Description: "test",
				Price:       100,
				Status:      orders.OrderStatus(1),
			})
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}
