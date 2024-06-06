package dto

import "time"

type ReportDTO struct {
	Id        int       `json:"id" binding:"required"`
	ProductId int       `json:"product_id" binding:"required"`
	ClientId  int       `json:"client_id"`
	Quantity  int       `json:"quantity"`
	Date      time.Time `json:"date"`
}
