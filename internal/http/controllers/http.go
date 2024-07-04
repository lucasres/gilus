package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasres/gilus/internal/domain/use_cases/crons"
)

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

}
