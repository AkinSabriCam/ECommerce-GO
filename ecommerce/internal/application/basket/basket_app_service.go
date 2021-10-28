package basket

import (
	"ecommerce/internal/application/basket/dtos"
	"ecommerce/internal/domain/basket"
	"github.com/google/uuid"
)

type AppService interface {
	AddNewProductToBasket(dto dtos.AddProductToBasketDto)
	GetMyBasket() dtos.BasketDto
}

type basketAppService struct {
	repository basket.IBasketRepository
	domainService basket.DomainService
}

func NewBasketAppService (repository basket.IBasketRepository, domainService basket.DomainService) AppService {
	return basketAppService{repository: repository, domainService: domainService}
}

func (service basketAppService) AddNewProductToBasket(dto dtos.AddProductToBasketDto) {
	// todo : userId will be get from token and pass to the domain service
	service.domainService.AddNewProductToBasket(dto.ProductId, dto.Quantity, uuid.New())
}

func (service basketAppService) GetMyBasket() dtos.BasketDto {
	// todo : will be implemented after domainservice implementation
	basket := service.repository.GetByUserId(uuid.New())
	return dtos.MapToBasketDto(basket)
}