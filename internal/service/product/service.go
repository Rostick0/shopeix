package product

import (
	"app/internal/domain/product"
	productRepository "app/internal/domain/product/repository"
	"app/internal/utils/pagination"
	"time"
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

func (s *Service) Update(input *product.UpdateProductRequest, productFinded *product.Product) (*product.Product, error) {
	if input.Title != nil {
		productFinded.Title = *input.Title
	}
	if input.OldPrice != nil {
		productFinded.OldPrice = *input.OldPrice
	}
	if input.Price != nil {
		productFinded.Price = *input.Price
	}
	if input.CategoryId != nil {
		productFinded.CategoryId = input.CategoryId
	}
	if input.IsShow != nil {
		productFinded.IsShow = *input.IsShow
	}

	productFinded.UpdatedAt = time.Now()

	return s.repo.Update(productFinded)
}

func (s *Service) FindAll(page int) (*[]product.Product, *pagination.Paginator, error) {
	return s.repo.FindAll(page)
}

func (s *Service) FindByID(id int64) (*product.Product, error) {
	return s.repo.FindByID(id)
}
