package productRepository

import (
	"app/internal/domain/product"
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

func (r *GormRepo) FindByID(id int64) (*product.Product, error) {
	var product product.Product
	if err := r.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return (&product), nil
}

func (r *GormRepo) FindAll(page int) (*[]product.Product, *pagination.Paginator, error) {
	var producs []product.Product

	p := &pagination.Paginator{Page: page, PerPage: 20}

	query := r.db.Model(&product.Product{}).Where("is_show = ?", true)

	_, err := p.Paginate(query, &producs)

	if err != nil {
		return nil, nil, err
	}

	return &producs, p, nil
}

func (r *GormRepo) Create(product *product.Product) (*product.Product, error) {
	r.db.Create(&product)

	return (product), nil
}

func (r *GormRepo) Update(product *product.Product) (*product.Product, error) {
	r.db.Save(&product)

	return (product), nil
}
