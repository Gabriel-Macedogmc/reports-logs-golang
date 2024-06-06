package interfaces

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
)

type ProdudctRepository interface {
	Create(data dto.ProductDTO) (*models.Product, error)
}
