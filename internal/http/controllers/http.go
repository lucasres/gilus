package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasres/gilus/internal/domain/entities"
	"github.com/lucasres/gilus/internal/domain/use_cases/crons"
)

type PingRequest struct {
	Name string `json:"name"`
}

func ListCronHandler(c *gin.Context) {
	uc := crons.NewListPingCronUseCase()

	pings, err := uc.Execute(c.Request.Context())
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, pings)
}

func PingCronHandler(c *gin.Context) {
	uc := crons.NewPingCronUseCase()
	request := PingRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	ping := entities.PingCron{
		Name:   request.Name,
		PingAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = uc.Execute(c.Request.Context(), ping)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
