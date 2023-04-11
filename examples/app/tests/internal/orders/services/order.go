package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"tests/internal/orders"
)

var ErrNotFound = errors.New("order is not found")

type OrderRepository interface {
	OrderGet(ctx context.Context, UUID uuid.UUID) (*orders.Order, error)
	OrderCreate(ctx context.Context, order *orders.Order) error
}

type OrderService struct {
	orderRepository OrderRepository
}

func NewOrderService(orderRepository OrderRepository) *OrderService {
	return &OrderService{orderRepository}
}

func (s *OrderService) OrderCreate(ctx context.Context, order *orders.Order) error {
	return s.orderRepository.OrderCreate(ctx, order)
}

func (s *OrderService) OrderGet(ctx context.Context, uuid uuid.UUID) (*orders.Order, error) {
	order, err := s.orderRepository.OrderGet(ctx, uuid)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return order, nil
}
