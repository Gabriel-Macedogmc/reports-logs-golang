package controllers

import (
	"net/http"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/queue"
	"github.com/gin-gonic/gin"
)

func GenerateReport(c *gin.Context) {
	var reportDto []dto.ReportDTO
	if err := c.ShouldBindJSON(&reportDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queue.Producer(reportDto)

	c.JSON(http.StatusOK, gin.H{})
}
