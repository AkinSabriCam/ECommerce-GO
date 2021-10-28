package category

type ICategoryRepository interface {
	GetAll() []Category
}
