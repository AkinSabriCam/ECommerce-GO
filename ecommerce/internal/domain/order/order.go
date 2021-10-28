package order

import "github.com/google/uuid"

type Order struct {
	ID uuid.UUID 			`gorm:"type:uuid;column:Id"`
	Status string  			`gorm:"column:Status"`
	Details []OrderDetail   `gorm:"foreignKey:OrderID"`
}

func NewOrder() Order {
	return Order{}
}

func (o *Order) GetId() uuid.UUID {
	return o.ID
}

func (o *Order) GetStatus() string {
	return o.Status
}

func (o *Order) GetDetails()  []OrderDetail {
	return o.Details
}

func (o *Order) SetId(id uuid.UUID) {
	o.ID = id
}

func (o *Order) SetStatus(status string) {
	o.Status = status
}

func (o *Order) AddDetails(details []OrderDetail) {
	o.Details = details
}