package dto

type ProductDTO struct {
	Name       string  `json:"name"`
	CategoryID uint    `json:"category_id"`
	Price      float64 `json:"price"`
}
