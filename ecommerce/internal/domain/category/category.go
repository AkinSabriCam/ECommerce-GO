package category

import (
	"ecommerce/internal/domain/product"
	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID 	`gorm:"type:uuid;column:Id"`
	Name string   	`gorm:"column:Name"`
	Products []product.Product  `gorm:"foreignKey:CategoryID"`
}

func NewCategory() Category {
	return Category{}
}

func (c *Category) GetId() uuid.UUID{
	return c.ID
}

func (c *Category) GetName() string{
	return c.Name
}
