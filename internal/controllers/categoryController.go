package controllers

import (
	"net/http"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/services"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (s *CategoryController) Create(c *gin.Context) {
	var categoryDto dto.CategoryDTO
	if err := c.ShouldBindJSON(&categoryDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := s.service.Create(categoryDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)

}
