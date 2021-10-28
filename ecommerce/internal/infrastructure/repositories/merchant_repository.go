package repositories

import (
	"ecommerce/internal/domain/merchant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository (db *gorm.DB) merchant.IMerchantRepository{
	return merchantRepository{db: db}
}

func (repository merchantRepository) GetById(id uuid.UUID) merchant.Merchant {
	var merchant merchant.Merchant
	repository.db.Where("Id = ?", id).Find(&merchant)
	return merchant
}

func (repository merchantRepository) GetAll() []merchant.Merchant {
	var merchantList []merchant.Merchant
	var merchants []merchant.Merchant
	repository.db.Find(&merchants)

	for _,merchant := range merchants {
		merchantList = append(merchantList, merchant)
	}

	return merchantList
}