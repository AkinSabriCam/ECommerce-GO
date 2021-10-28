package order

import (
	"ecommerce/internal/application/order/dtos"
	"ecommerce/internal/domain/order"
	"github.com/google/uuid"
)

type AppService interface {
	GetAll() []dtos.OrderDto
	GetById(id uuid.UUID) dtos.OrderDto
	Create(dto dtos.CreateOrderDto) dtos.OrderDto
	Cancel(id uuid.UUID)
}

type orderAppService struct {
	repository order.IOrderRepository
	domainService order.DomainService
}

func NewOrderAppService (repository order.IOrderRepository, domainService order.DomainService) AppService {
	return orderAppService{
		repository: repository,
		domainService: domainService,
	}
}

func (appService  orderAppService) GetAll() []dtos.OrderDto{
	orders := appService.repository.GetAll()
	return dtos.MapToOrderDtoSlice(orders)
}

func (appService  orderAppService) GetById(id uuid.UUID) dtos.OrderDto{
	order := appService.repository.GetById(id)
	return dtos.MapToOrderDto(order)
}

func (appService  orderAppService) Create(dto dtos.CreateOrderDto) dtos.OrderDto{
	order := appService.domainService.Create(dto.Status,dto.ProductID, dto.Quantity, dto.Amount)
	return dtos.MapToOrderDto(order)
}

func (appService  orderAppService) Cancel(id uuid.UUID){
	appService.domainService.Cancel(id)
}