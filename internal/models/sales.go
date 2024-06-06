package models

import (
	"time"

	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model
	ID        uint
	ProductId uint `gorm:"product_id"`
	ClientId  uint `gorm:"client_id"`
	Quantity  uint `gorm:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
