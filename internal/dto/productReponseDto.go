package dto

import "time"

type ProductResponseDTO struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	CategoryID uint      `json:"category_id"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
