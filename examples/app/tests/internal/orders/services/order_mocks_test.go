//go:build unit_test

package services

import (
	"context"
	"github.com/google/uuid"
	"tests/internal/orders"
)

// OrderRepositoryMock приклад опису мок даних
// - Для кожного методу всі його значення, що повертаються мають бути описані як поля структури.
// - Кожен метод повертає лише ці поля, без будь якої логіки.
// Таким чином дуже легко можна для кожного тесту описувати конкретні очікувані значення.
type OrderRepositoryMock struct {
	GetSuccess  *orders.Order
	GetError    error
	CreateError error
}

func (r *OrderRepositoryMock) OrderGet(ctx context.Context, uuid uuid.UUID) (*orders.Order, error) {
	return r.GetSuccess, r.GetError
}

func (r *OrderRepositoryMock) OrderCreate(ctx context.Context, order *orders.Order) error {
	return r.CreateError
}
