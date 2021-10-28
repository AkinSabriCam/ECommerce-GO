package merchant

import (
	"ecommerce/internal/domain/product"
	"github.com/google/uuid"
)

type Merchant struct{
	ID uuid.UUID 			`gorm:"type:uuid;column:Id"`
	BusinessName string 	`gorm:"column:BusinessName"`
	Contact MerchantContact `gorm:"foreignKey:MerchantID"`
	Products  []product.Product `gorm:"foreignKey:MerchantID"`

}

func NewMerchant() Merchant {
	return  Merchant{}
}

func (m *Merchant) GetId() uuid.UUID{
	return m.ID
}

func (m *Merchant) GetBusinessName() string{
	return m.BusinessName
}

func (m *Merchant) GetMerchantContact() MerchantContact{
	return m.Contact
}

