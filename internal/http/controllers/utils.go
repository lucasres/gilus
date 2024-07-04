package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, err error) {
	fmt.Printf("erro when execute list ping: %s\n", err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
}
