package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "healthy",
	})
}

func HandleLogin(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"data": "logged-in",
	})
}

