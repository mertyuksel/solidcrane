package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type text struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.POST("/uppercase", convertUpperCase)
	router.Run("0.0.0.0:3000")
}

func convertUpperCase(c *gin.Context) {
	var newText text

	if err := c.BindJSON(&newText); err != nil {
		return
	}

	newText.Message = strings.ToUpper(newText.Message)

	c.JSON(http.StatusOK, newText)
}
