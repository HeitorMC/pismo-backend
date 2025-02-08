package main

import (
	"net/http"

	"github.com/HeitorMC/pismo-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {

	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
