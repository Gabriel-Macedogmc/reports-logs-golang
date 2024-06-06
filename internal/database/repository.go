package database

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
	"gorm.io/gorm"
)

type SalesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) *SalesRepository {
	return &SalesRepository{db: db}
}

func (r *SalesRepository) GetAll() ([]models.Sales, error) {
	var sales []models.Sales

	err := r.db.Find(&sales).Error

	if err != nil {
		return nil, err
	}

	return sales, nil
}
