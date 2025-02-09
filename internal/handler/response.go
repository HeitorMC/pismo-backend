package handler

import (
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendResponse(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	if msg != "" {
		ctx.JSON(code, gin.H{
			"message": msg,
		})
	} else {
		ctx.JSON(code, data)
	}
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
