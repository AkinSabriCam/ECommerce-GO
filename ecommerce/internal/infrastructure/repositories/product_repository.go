package repositories

import (
	"ecommerce/internal/domain/product"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.IProductRepository {
	return ProductRepository{db: db}
}

func (repo ProductRepository) GetById(id uuid.UUID) product.Product {
	var entity = product.Product{}
	repo.db.Where(product.Product{ID: id}).Find(&entity)
	return entity
}

func (repo ProductRepository) GetAllByCategoryId(categoryId uuid.UUID) []product.Product {
	var products []product.Product
	repo.db.Where(product.Product{CategoryID: categoryId}).Find(&products)
	return products
}

func (repo ProductRepository) GetAllByMerchantId(merchantId uuid.UUID) []product.Product {
	var products []product.Product
	repo.db.Where(product.Product{MerchantID:merchantId}).Find(&products)

	return products
}

func (repo ProductRepository) GetAll() []product.Product {
	var products []product.Product
	repo.db.Find(&products)
	return products
}
