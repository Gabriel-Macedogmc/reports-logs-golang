package interfaces

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
)

type CategoryRepository interface {
	Create(data dto.CategoryDTO) (*models.Category, error)
}
