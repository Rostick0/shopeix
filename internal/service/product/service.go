package product

import (
	"app/internal/domain/product"
	productRepository "app/internal/domain/product/repository"
)

type Service struct {
	repo *productRepository.GormRepo
}

func NewService(repo *productRepository.GormRepo) *Service {
	return &Service{repo: repo}
}

// func (s *Service) FindByID(id int64) (*category.Category, error) {
// 	return s.repo.FindByID(id)
// }

func (s *Service) Create(input *product.CreateProductRequest) (*product.Product, error) {
	product := &product.Product{}

	return s.repo.Create(product)
}
