package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         uint
	Name       string
	Price      float64
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
