package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Gabriel-Macedogmc/report-logs-golang/di"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/controllers"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/database"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/middleware"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/queue"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()

	r := gin.Default()

	clientController := di.ClientDI(db)

	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/client", clientController.Create)
		publicRoutes.POST("/client/login", clientController.Login)

		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.POST("/report", controllers.GenerateReport)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		queue.Consumer(signals)
	}()

	go func() {
		r.Run(":3000")
	}()

	<-signals
	log.Println("Shutting down")
}
