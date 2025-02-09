package router

import (
	docs "github.com/HeitorMC/pismo-backend/docs"
	"github.com/HeitorMC/pismo-backend/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	router.POST("/accounts", handler.CreateAccount)
	router.GET("/accounts/:accountId", handler.ShowAccount)
	router.POST("/transactions", handler.CreateTransaction)

	docs.SwaggerInfo.Title = "Pismo Backend"
	docs.SwaggerInfo.Description = "This is a simple API to create accounts and transactions"
	docs.SwaggerInfo.Version = "1.0"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
