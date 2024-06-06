package repositories

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &repository{db: db}
}

// Create implements interfaces.CategoryRepository.
func (r *repository) Create(data dto.CategoryDTO) (*models.Category, error) {
	category := models.Category{
		Name: data.Name,
	}

	if err := r.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
