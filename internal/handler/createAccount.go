package handler

import (
	"net/http"

	"github.com/HeitorMC/pismo-backend/internal/models"
	"github.com/HeitorMC/pismo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateAccount(ctx *gin.Context) {
	request := CreateAccountRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	account := models.Account{
		DocumentNumber: request.DocumentNumber,
	}

	if err := service.CreateAccount(&account); err != nil {
		sendError(ctx, http.StatusForbidden, err.Error())
		return
	}

	sendResponse(ctx, http.StatusCreated, "Account successfully created", nil)
}
