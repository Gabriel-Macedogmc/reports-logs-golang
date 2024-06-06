package di

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/controllers"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/database/repositories"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/services"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/utils"
	"gorm.io/gorm"
)

func ClientDI(db *gorm.DB) *controllers.ClientController {
	repository := repositories.NewClientRepository(db)
	jwtLogin := utils.NewJwtLogin()
	bcrypt := utils.NewBcrypt()
	service := services.NewClientsService(repository, jwtLogin, bcrypt)
	controller := controllers.NewClientController(*service)

	return controller
}
