package order

import (
	"github.com/google/uuid"
)

type OrderDetail struct {
	ID uuid.UUID  			`gorm:"type:uuid;column:Id"`
	OrderID uuid.UUID  		`gorm:"type:uuid;column:OrderId"`
	ProductID uuid.UUID 	`gorm:"type:uuid;column:ProductId"`
	Quantity int 			`gorm:"column:Quantity"`
	Amount float64 			`gorm:"column:Amount"`
	Date uuid.Time 			`gorm:"column:Date"`
}

func NewOrderDetail() OrderDetail{
	return OrderDetail{}
}

func (od * OrderDetail) GetId() uuid.UUID{
	return od.ID
}

func (od * OrderDetail) GetOrderId() uuid.UUID{
	return od.OrderID
}

func (od * OrderDetail) GetProductId() uuid.UUID{
	return od.ProductID
}

func (od * OrderDetail) GetQuantity() int{
	return od.Quantity
}

func (od * OrderDetail) GetAmount() float64{
	return od.Amount
}

func (od * OrderDetail) GetDate() uuid.Time{
	return od.Date
}

// set methods

func (od * OrderDetail) SetId(id uuid.UUID) {
	od.ID = id
}

func (od * OrderDetail) SetOrderId(orderId uuid.UUID) {
	od.OrderID= orderId
}

func (od * OrderDetail) SetProductId(productId uuid.UUID) {
	od.ProductID = productId
}

func (od * OrderDetail) SetQuantity(quantity int) {
	od.Quantity = quantity
}

func (od * OrderDetail) SetAmount(amount float64) {
	od.Amount = amount
}

func (od * OrderDetail) SetDate(date uuid.Time) {
	od.Date = date
}