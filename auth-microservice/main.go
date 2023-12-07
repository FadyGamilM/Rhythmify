package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var api_port = 5003

func init() {
	api_port, _ = strconv.Atoi(os.Getenv("PORT"))
}

func main() {
	router := gin.Default()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", api_port),
		Handler: router,
	}
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "healthy",
		})
	})
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
}
