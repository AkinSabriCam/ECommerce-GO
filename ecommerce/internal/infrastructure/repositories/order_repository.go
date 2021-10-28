package repositories

import (
	"ecommerce/internal/domain/order"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository (db *gorm.DB) order.IOrderRepository {
	return orderRepository{db: db}
}

func (repository orderRepository) Add(entity order.Order) order.Order{
	repository.db.Model(&order.Order{}).Create(entity)
	return entity
}

func (repository orderRepository) GetById(id uuid.UUID) order.Order{
	var entity order.Order
	repository.db.Where("Id = ?", id).Find(&entity)
	return entity
}

func (repository orderRepository) GetAll() []order.Order{
	var orders []order.Order
	repository.db.Find(&orders)

	var orderList []order.Order
	for _,order := range orders {
		orderList = append(orderList, order)
	}
	return orderList
}

func (repository orderRepository) Update(entity order.Order) order.Order{
	repository.db.Save(&entity)
	return entity
}

func (repository orderRepository) Delete(id uuid.UUID) {
	var entity order.Order
	repository.db.Find(&entity)
	repository.db.Delete(entity)
}
