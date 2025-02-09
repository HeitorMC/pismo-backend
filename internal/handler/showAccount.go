package handler

import (
	"net/http"

	"github.com/HeitorMC/pismo-backend/internal/models"
	"github.com/HeitorMC/pismo-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func ShowAccount(ctx *gin.Context) {
	id := ctx.Param("accountId")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	account := models.Account{}

	if err := service.FindAccount(&account, id); err != nil {
		sendError(ctx, http.StatusNotFound, "account not found")
		return
	}

	sendResponse(ctx, http.StatusOK, "", account)
}
