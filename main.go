package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasres/gilus/internal/http/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	cron := router.Group("/crons")
	cron.GET("/", controllers.ListCronHandler)
	cron.POST("/", controllers.PingCronHandler)

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("cant listiner in 8000: %e", err)
	}

}
