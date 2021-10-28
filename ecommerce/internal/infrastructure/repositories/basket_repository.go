package repositories

import (
	"ecommerce/internal/domain/basket"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type basketRepository struct {
	db *gorm.DB
}

func NewBasketRepository(db *gorm.DB) basket.IBasketRepository {
	return basketRepository{db : db}
}

func (repo basketRepository) Add(entity basket.Basket) basket.Basket{
	repo.db.Create(&entity)
	return entity
}

func (repo basketRepository) GetByUserId(userId uuid.UUID) basket.Basket{
	var basket basket.Basket
	repo.db.Where("UserId = ?", userId).First(&basket)
	return basket
}

func (repo basketRepository) Update(entity basket.Basket) basket.Basket{
	repo.db.Save(&entity)
	return entity
}

func (repo basketRepository) Delete(id uuid.UUID){
	var basket basket.Basket
	repo.db.Where("Id = ?", id).First(&basket)
	repo.db.Delete(&basket)
}