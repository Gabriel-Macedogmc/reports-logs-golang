package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
