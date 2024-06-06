package services

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
)

type CategoryService struct {
	repo interfaces.CategoryRepository
}

func NewCategoryService(repo interfaces.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(data dto.CategoryDTO) (*models.Category, error) {
	category, err := s.repo.Create(data)

	if err != nil {
		return nil, err
	}

	return category, nil
}
