package merchant

import (
	"ecommerce/internal/application/merchant/dtos"
	"ecommerce/internal/domain/merchant"
	"github.com/google/uuid"
)

type AppService interface {
	GetById(id uuid.UUID) dtos.MerchantDto
	GetAll() [] dtos.MerchantDto
}

type merchantAppService struct {
	repository merchant.IMerchantRepository
}

func NewMerchantAppService (repository merchant.IMerchantRepository) AppService {
	return merchantAppService{repository: repository}
}

func (service merchantAppService) GetById(id uuid.UUID) dtos.MerchantDto {
	var merchant = service.repository.GetById(id)
	return dtos.MapToMerchantDto(merchant)
}

func (service merchantAppService) GetAll() []dtos.MerchantDto {
	var merchants = service.repository.GetAll()
	return dtos.MapToMerchantSliceDto(merchants)
}