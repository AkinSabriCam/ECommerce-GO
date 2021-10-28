package product

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID  			`gorm:"type:uuid;primaryKey;column:Id"`
	CategoryID uuid.UUID 	`gorm:"type:uuid;column:CategoryId"`
	MerchantID uuid.UUID 		`gorm:"type:uuid;column:MerchantId"`
	Brand string 			`gorm:"column:Brand"`
	Model string 			`gorm:"column:Model"`
	Description string 		`gorm:"column:Description"`
	Quantity int 			`gorm:"column:Quantity"`
	Amount float64			`gorm:"column:Amount"`
}

func NewProduct() Product {
	return Product{}
}

func (p *Product) GetId() uuid.UUID{
	return p.ID
}

func (p *Product) GetCategoryId() uuid.UUID{
	return p.CategoryID
}

func (p *Product) GetBrand() string{
	return p.Brand
}

func (p *Product) GetModel() string{
	return p.Model
}

func (p *Product) GetDescription() string{
	return p.Description
}

func (p *Product) GetQuantity() int{
	return p.Quantity
}

func (p *Product) GetAmount() float64{
	return p.Amount
}