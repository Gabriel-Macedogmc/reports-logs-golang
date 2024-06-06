package database

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:docker-mysql@tcp(localhost:3306)/reports"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	models := []any{models.Sales{}, models.Client{}, models.Category{}, models.Product{}}

	autoMigrate(models, db)

	return db
}

func autoMigrate(models []any, db *gorm.DB) {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
}
