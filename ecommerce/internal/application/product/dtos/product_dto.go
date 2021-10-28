package dtos

import (
	"ecommerce/internal/domain/product"
	"github.com/google/uuid"
)

type ProductDto struct {
	Id uuid.UUID
	Brand string
	Model string
	CategoryId uuid.UUID
	Description string
	Stock int
}

func MapToProductDtoSlice (products []product.Product) []ProductDto {
	var productList []ProductDto

	for _,product := range products {
		productList = append(productList, MapToProductDto(product))
	}

	return productList
}

func MapToProductDto (product product.Product) ProductDto {
	return ProductDto{
		Id: product.GetId(),
		Brand: product.GetBrand(),
		Model: product.GetModel(),
		CategoryId: product.GetCategoryId(),
		Description: product.GetDescription(),
		Stock: product.GetQuantity(),
	}
}