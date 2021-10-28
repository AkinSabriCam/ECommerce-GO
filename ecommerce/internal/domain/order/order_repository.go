package order

import "github.com/google/uuid"

type IOrderRepository interface {
	Add(entity Order) Order
	GetById(id uuid.UUID) Order
	GetAll() []Order
	Update(entity Order) Order
	Delete(id uuid.UUID)
}
