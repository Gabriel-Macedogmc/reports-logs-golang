package interfaces

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
)

type ClientRepository interface {
	Create(data dto.CreateClient) (*models.Client, error)
	GetByEmail(email string) *models.Client
}
