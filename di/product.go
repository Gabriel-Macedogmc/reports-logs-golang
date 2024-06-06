package di

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/controllers"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/database/repositories"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/services"
	"gorm.io/gorm"
)

func ProductDI(db *gorm.DB) *controllers.ProductController {
	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)
	controller := controllers.NewProductController(*service)

	return controller
}
