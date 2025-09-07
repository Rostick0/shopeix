package product

import (
	"app/internal/domain/product"
	productRepository "app/internal/domain/product/repository"
	"app/internal/utils/pagination"
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
	product := &product.Product{
		Title:      input.Title,
		OldPrice:   input.OldPrice,
		Price:      input.Price,
		CategoryId: input.CategoryId,
		IsShow:     input.IsShow,
	}

	return s.repo.Create(product)
}

func (s *Service) FindAll(page int) (*[]product.Product, *pagination.Paginator, error) {
	return s.repo.FindAll(page)
}
