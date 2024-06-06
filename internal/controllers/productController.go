package controllers

import (
	"net/http"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (s *ProductController) Create(c *gin.Context) {
	var productDto dto.ProductDTO
	if err := c.ShouldBindJSON(&productDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := s.service.Create(productDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
