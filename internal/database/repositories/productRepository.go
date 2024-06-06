package repositories

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProdudctRepository {
	return &productRepository{db: db}
}

func (repo *productRepository) Create(data dto.ProductDTO) (*models.Product, error) {
	product := models.Product{
		Name:       data.Name,
		Price:      data.Price,
		CategoryID: data.CategoryID,
	}

	if err := repo.db.Omit("Category").Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
