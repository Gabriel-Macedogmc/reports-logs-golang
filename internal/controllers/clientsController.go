package controllers

import (
	"net/http"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/services"
	"github.com/gin-gonic/gin"
)

type ClientController struct {
	service services.ClientsService
}

func NewClientController(service services.ClientsService) *ClientController {
	return &ClientController{service: service}
}

func (s *ClientController) Create(c *gin.Context) {
	var clientData dto.CreateClient
	if err := c.ShouldBindJSON(&clientData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := s.service.Create(clientData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func (s *ClientController) Login(c *gin.Context) {
	var loginData dto.CreateLoginDTO
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.service.Login(loginData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
