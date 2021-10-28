package basket

import (
	"ecommerce/internal/domain/product"
	"github.com/google/uuid"
	"time"
)

type DomainService interface {
	AddNewProductToBasket(productId uuid.UUID, quantity int, userId uuid.UUID)
}

type domainService struct {
	repo IBasketRepository
	productRepo product.IProductRepository
}

func NewBasketDomainService(repo IBasketRepository, productRepo product.IProductRepository) DomainService {
	return domainService{repo : repo, productRepo: productRepo}
}

func (dS domainService) AddNewProductToBasket(productId uuid.UUID, quantity int, userId uuid.UUID){
	product :=  dS.productRepo.GetById(productId)
	basket := dS.repo.GetByUserId(userId)
	detail := NewBasketDetail()

	detail.SetQuantity(quantity)
	detail.SetProductId(productId)
	detail.SetAmount(product.GetAmount() * float64(quantity))
	detail.SetDate(time.Now())

	basket.AddDetail(detail)
	dS.repo.Update(basket)
}