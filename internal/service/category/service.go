package category

import (
	"app/internal/domain/category"
	categoryRepository "app/internal/domain/category/repository"
	"app/internal/utils/pagination"
)

type Service struct {
	repo *categoryRepository.CategoryGormRepo
}

func NewService(repo *categoryRepository.CategoryGormRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindByID(id int64) (*category.Category, error) {
	return s.repo.FindByID(id)
}

func (s *Service) FindAll(page int) (*[]category.Category, *pagination.Paginator, error) {
	return s.repo.FindAll(page)
}
