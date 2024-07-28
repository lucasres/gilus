package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasres/gilus/internal/domain/entities"
	"github.com/lucasres/gilus/internal/domain/use_cases/crons"
)

type PingRequest struct {
	Name string `json:"name"`
}

func ListCronHandler(c *gin.Context) {
	uc := crons.NewListCronUseCase()

	pings, err := uc.Execute(c.Request.Context())
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, pings)
}

func PingCronHandler(c *gin.Context) {
	request := PingRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	ping := entities.PingCron{
		Name: request.Name,
	}

	uc := crons.NewPingCronUseCase()
	err := uc.Execute(c.Request.Context(), ping)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func ListPingCronHandler(c *gin.Context) {
	uc := crons.NewListPingCronUseCase()
	pings, err := uc.Execute(c.Request.Context(), c.Param("name"))
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, pings)
}
