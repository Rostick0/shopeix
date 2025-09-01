package categoryRepository

import (
	"app/internal/domain/category"
	"app/internal/utils/pagination"
	"errors"

	"gorm.io/gorm"
)

type GormRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) *GormRepo {
	return &GormRepo{db: db}
}

func (r *GormRepo) FindByID(id int64) (*category.Category, error) {
	var category category.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // пользователь не найден
		}
		return nil, err // другая ошибка
	}
	return (&category), nil
}

func (r *GormRepo) FindAll(page int) (*[]category.Category, *pagination.Paginator, error) {
	var categories []category.Category

	p := &pagination.Paginator{Page: page, PerPage: 20}

	query := r.db.Model(&category.Category{})

	_, err := p.Paginate(query, &categories)

	if err != nil {
		return nil, nil, err
	}

	return &categories, p, nil
}
