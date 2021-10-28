package category

import "ecommerce/internal/domain/category"

type AppService interface {
	GetAll() []category.Category
}

type appService struct {
	repo category.ICategoryRepository
}

func NewAppService(repo category.ICategoryRepository) AppService{
	return appService{repo: repo}
}

func (service appService) GetAll() []category.Category{
	return service.repo.GetAll()
}