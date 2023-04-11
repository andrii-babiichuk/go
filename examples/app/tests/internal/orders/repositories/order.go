package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"tests/internal/orders"
)

const selectByIdQuery = "SELECT uuid,date,description,price,status FROM orders WHERE uuid=$1"
const insertQuery = "INSERT INTO orders (uuid,date,description,price,status) VALUES ($1, $2, $3, $4, $5)"

type OrderRepository struct {
	c *pgxpool.Pool
}

func NewOrderRepository(c *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{c}
}

func (o *OrderRepository) OrderGet(ctx context.Context, uuid uuid.UUID) (*orders.Order, error) {
	var entity orders.Order
	row := o.c.QueryRow(ctx, selectByIdQuery, uuid.String())
	err := row.Scan(
		&entity.UUID,
		&entity.Date,
		&entity.Description,
		&entity.Price,
		&entity.Status,
	)
	return &entity, err
}

func (o *OrderRepository) OrderCreate(ctx context.Context, order *orders.Order) error {
	_, err := o.c.Exec(ctx, insertQuery,
		order.UUID.String(),
		order.Date,
		order.Description,
		order.Price,
		order.Status,
	)
	return err
}
