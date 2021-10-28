package order

import "github.com/google/uuid"

type DomainService interface {
	Create(status string, productId uuid.UUID, quantity int, amount float64) Order
	Cancel(id uuid.UUID)
}

type domainService struct {
	repository IOrderRepository
}

func NewDomainService (repository IOrderRepository) DomainService {
	return domainService{repository: repository}
}

func (dS domainService) Create(status string, productId uuid.UUID, quantity int, amount float64) Order {
	var order = NewOrder()
	var detail = NewOrderDetail()

	detail.SetId(uuid.New())
	detail.SetAmount(amount)
	detail.SetQuantity(quantity)
	detail.SetProductId(productId)

	order.SetId(uuid.New())
	order.SetStatus(status)
	order.AddDetails([]OrderDetail{detail})
	dS.repository.Add(order)

	return order
}

func (dS domainService) Cancel(id uuid.UUID) {
	var order = dS.repository.GetById(id)
	order.SetStatus("Canceled")
	dS.repository.Update(order)
}