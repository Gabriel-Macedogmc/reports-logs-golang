package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        uint
	Name      string
	Email     string `gorm:"uniqueIndex, length:50"`
	Document  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
