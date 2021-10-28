package basket

import (
	"github.com/google/uuid"
	"time"
)

type BasketDetail struct {
	ID uuid.UUID `gorm:"type:uuid;column:Id"`
	UserID uuid.UUID `gorm:"type:uuid;column:UserId"`
	BasketID uuid.UUID `gorm:"type:uuid;column:BasketId"`
	ProductID uuid.UUID `gorm:"type:uuid;column:ProductId"`
	Quantity int `gorm:"type:int;column:Quantity"`
	Amount float64 `gorm:"type:numeric;column:Amount"`
	Date time.Time `gorm:"type:date;column:Date"`
}

func NewBasketDetail() BasketDetail{
	return BasketDetail{}
}

func (od * BasketDetail) GetId() uuid.UUID{
	return od.ID
}

func (od * BasketDetail) GetBasketId() uuid.UUID{
	return od.BasketID
}

func (od * BasketDetail) GetProductId() uuid.UUID{
	return od.ProductID
}

func (od * BasketDetail) GetQuantity() int{
	return od.Quantity
}

func (od * BasketDetail) GetAmount() float64{
	return od.Amount
}

func (od * BasketDetail) GetDate() time.Time{
	return od.Date
}

// set methods

func (od * BasketDetail) SetId(id uuid.UUID) {
	od.ID = id
}

func (od * BasketDetail) SetBasketId(basketId uuid.UUID) {
	od.BasketID= basketId
}

func (od * BasketDetail) SetProductId(productId uuid.UUID) {
	od.ProductID = productId
}

func (od * BasketDetail) SetQuantity(quantity int) {
	od.Quantity = quantity
}

func (od * BasketDetail) SetAmount(amount float64) {
	od.Amount = amount
}

func (od * BasketDetail) SetDate(date time.Time) {
	od.Date = date
}

