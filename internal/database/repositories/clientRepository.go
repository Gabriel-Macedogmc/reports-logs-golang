package repositories

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) interfaces.ClientRepository {
	return &clientRepository{db: db}
}

func (c *clientRepository) Create(data dto.CreateClient) (*models.Client, error) {
	user := models.Client{
		Name:     data.Name,
		Email:    data.Email,
		Document: data.Document,
	}

	c.db.Create(&user)

	return &user, nil
}

// GetByEmail implements interfaces.ClientRepository.
func (c *clientRepository) GetByEmail(email string) *models.Client {
	client := models.Client{}

	c.db.Select("id", "name", "email").Where("email = ?", email).Find(&client)
	return &client
}
