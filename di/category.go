package di

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/controllers"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/database/repositories"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/services"
	"gorm.io/gorm"
)

func CategoryDI(db *gorm.DB) *controllers.CategoryController {
	repository := repositories.NewRepository(db)
	service := services.NewCategoryService(repository)
	controller := controllers.NewCategoryController(*service)

	return controller
}
