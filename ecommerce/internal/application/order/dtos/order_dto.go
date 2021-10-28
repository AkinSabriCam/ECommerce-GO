package dtos

import (
	"ecommerce/internal/domain/order"
	"github.com/google/uuid"
)

type OrderDto struct {
	Id uuid.UUID
	Status string
	Details []order.OrderDetail
}

type CreateOrderDto struct {
	Status string
	ProductID uuid.UUID
	Quantity int
	Amount float64
	Date uuid.Time
}

func MapToOrderDto(order order.Order) OrderDto {
	return OrderDto{
		Id: order.GetId(),
		Status: order.GetStatus(),
		Details: order.GetDetails(),
	}
}

func MapToOrderDtoSlice(orderSlice []order.Order) []OrderDto {
	var orderDtoSlice []OrderDto
	for _,order := range orderSlice {
		orderDtoSlice = append(orderDtoSlice, MapToOrderDto(order))
	}
	return orderDtoSlice
}
