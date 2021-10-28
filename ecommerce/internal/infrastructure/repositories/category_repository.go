package repositories

import (
	"ecommerce/internal/domain/category"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) category.ICategoryRepository {
	return categoryRepository{db : db}
}

func (repository categoryRepository) GetAll() []category.Category {
	var categories []category.Category
	repository.db.Find(&categories)

	return categories
}
