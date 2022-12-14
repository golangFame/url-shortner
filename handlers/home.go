package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HomeHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Service",
	})
}
