package product

import (
	"ecommerce/internal/application/product/dtos"
	"ecommerce/internal/domain/product"
	"github.com/google/uuid"
)

type AppService interface{
	GetAll() []dtos.ProductDto
	GetAllByMerchantId(merchantId uuid.UUID) []dtos.ProductDto
	GetById(id uuid.UUID) dtos.ProductDto
}

type productAppService struct {
	repository product.IProductRepository
}

func NewProductAppService (repository product.IProductRepository) AppService {
	return productAppService{repository: repository}
}

func (appService productAppService) GetAll() []dtos.ProductDto {
	var products = appService.repository.GetAll()
	return dtos.MapToProductDtoSlice(products)
}

func (appService productAppService) GetAllByMerchantId(merchantId uuid.UUID) []dtos.ProductDto {
	var products = appService.repository.GetAllByMerchantId(merchantId)
	return dtos.MapToProductDtoSlice(products)
}

func (appService productAppService) GetById(id uuid.UUID) dtos.ProductDto{
	var product = appService.repository.GetById(id)
	return dtos.MapToProductDto(product)
}