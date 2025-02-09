package router

import (
	"github.com/HeitorMC/pismo-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	router.POST("/accounts", handler.CreateAccount)
	router.GET("/accounts/:accountId", handler.ShowAccount)
}
