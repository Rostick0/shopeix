package categoryRepository

import (
	"app/internal/domain/category"
	"app/internal/utils/pagination"
	"errors"

	"gorm.io/gorm"
)

type CategoryGormRepo struct {
	db *gorm.DB
}

func NewCategoryGormRepo(db *gorm.DB) *CategoryGormRepo {
	return &CategoryGormRepo{db: db}
}

func (r *CategoryGormRepo) FindByID(id int64) (*category.Category, error) {
	var category category.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // пользователь не найден
		}
		return nil, err // другая ошибка
	}
	return (&category), nil
}

func (r *CategoryGormRepo) FindAll(page int) (*[]category.Category, *pagination.Paginator, error) {
	var categories []category.Category

	// if err := r.db.Find(&categories).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return nil, nil // пользователь не найден
	// 	}
	// 	return nil, err // другая ошибка
	// }

	// r.db.Model()

	p := &pagination.Paginator{Page: page, PerPage: 20}

	query := r.db.Model(&category.Category{})

	_, err := p.Paginate(query, &categories)

	if err != nil {
		return nil, nil, err
	}

	return &categories, p, nil
}

// func (r *CategoryGormRepo) Create(u *category.Category) error {
// 	// реализация через GORM
// }
