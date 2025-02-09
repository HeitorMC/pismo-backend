package handler

import (
	"net/http"

	"github.com/HeitorMC/pismo-backend/internal/models"
	"github.com/HeitorMC/pismo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Show an account
// @Description Find an account by its ID
// @Tags Accounts
// @Accept json
// @Produce json
// @Param accountId path string false "Account ID"
// @Success 200 {object} ShowAccountResponse
// @Failure 404 {object} ErrorResponse
// @Router /accounts/{accountId} [get]
func ShowAccount(ctx *gin.Context) {
	id := ctx.Param("accountId")

	account := models.Account{}

	if err := service.FindAccount(&account, id); err != nil {
		sendError(ctx, http.StatusNotFound, "account not found")
		return
	}

	sendResponse(ctx, http.StatusOK, "", account)
}
