package basket

import "github.com/google/uuid"

type Basket struct {
	ID uuid.UUID `gorm:"type:uuid;column:Id"`
	Details []BasketDetail `gorm:"foreignKey:BasketID"`
}

func NewBasket() Basket {
	return Basket{}
}

func (o *Basket) GetId() uuid.UUID {
	return o.ID
}

func (o *Basket) GetDetails()  []BasketDetail {
	return o.Details
}

func (o *Basket) SetId(id uuid.UUID) {
	o.ID = id
}

func (o *Basket) AddDetail(detail BasketDetail) {
	o.Details = append(o.Details, detail)
}

