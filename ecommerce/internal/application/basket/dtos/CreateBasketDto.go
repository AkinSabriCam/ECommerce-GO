package dtos

import (
	"ecommerce/internal/domain/basket"
	"github.com/google/uuid"
)

type AddProductToBasketDto struct {
	ProductId uuid.UUID
	Quantity int
}

type BasketDto struct {
	Id uuid.UUID
	TotalAmount uuid.UUID
	BasketDetails []basket.BasketDetail
}

func MapToBasketDto(basket basket.Basket) BasketDto {
	return BasketDto{
		Id : basket.GetId(),
		BasketDetails: basket.GetDetails(),
	}
}