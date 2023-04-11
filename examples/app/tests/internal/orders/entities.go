package orders

import (
	"github.com/google/uuid"
	"time"
)

type OrderStatus int8

const (
	Init OrderStatus = iota
	Completed
	Canceled
)

type Order struct {
	UUID        uuid.UUID   `json:"uuid,omitempty"`
	Date        time.Time   `json:"date"`
	Description string      `json:"description,omitempty"`
	Price       int64       `json:"price,omitempty"`
	Status      OrderStatus `json:"status,omitempty"`
}
