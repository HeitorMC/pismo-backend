package handler

import (
	"net/http"

	"github.com/HeitorMC/pismo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Create a transaction
// @Description Create a new transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body CreateTransactionRequest true "Request body"
// @Success 201 {object} CreateTransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Router /transactions [post]
func CreateTransaction(ctx *gin.Context) {
	request := CreateTransactionRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transactionData := service.CreateTransactionData{
		AccountID:       *request.AccountID,
		OperationTypeID: *request.OperationTypeID,
		Amount:          *request.Amount,
	}

	if err := service.CreateTransaction(&transactionData); err != nil {
		sendError(ctx, http.StatusForbidden, err.Error())
		return
	}

	sendResponse(ctx, http.StatusCreated, "Transaction successfully created", nil)
}
