package services

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
)

type ProductService struct {
	repo interfaces.ProdudctRepository
}

func NewProductService(repo interfaces.ProdudctRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(data dto.ProductDTO) (*dto.ProductResponseDTO, error) {
	product, err := s.repo.Create(data)

	if err != nil {
		return nil, err
	}

	return &dto.ProductResponseDTO{
		ID:         product.ID,
		Name:       product.Name,
		CategoryID: product.CategoryID,
		Price:      product.Price,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}, nil
}
