package handler

import (
	"net/http"

	"github.com/HeitorMC/pismo-backend/internal/models"
	"github.com/HeitorMC/pismo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Create an account
// @Description Create a new account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param request body CreateAccountRequest true "Request body (DocumentNumber must be an 11 character long string)"
// @Success 201 {object} CreateAccountResponse
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Router /accounts [post]
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
