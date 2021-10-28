package basket

import "github.com/google/uuid"

type IBasketRepository interface {
	Add(entity Basket) Basket
	GetByUserId(userId uuid.UUID) Basket
	Update(entity Basket) Basket
	Delete(id uuid.UUID)
}
