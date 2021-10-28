package product

import "github.com/google/uuid"

type IProductRepository interface {
	GetById(id uuid.UUID) Product
	GetAllByCategoryId(categoryId uuid.UUID) []Product
	GetAllByMerchantId(merchantId uuid.UUID) []Product
	GetAll() []Product
}
