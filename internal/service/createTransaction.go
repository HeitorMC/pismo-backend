package service

import (
	"fmt"
	"math"

	"github.com/HeitorMC/pismo-backend/config"
	"github.com/HeitorMC/pismo-backend/internal/models"
	"github.com/HeitorMC/pismo-backend/internal/repository"
)

type CreateTransactionData struct {
	AccountID       uint    `json:"account_id"`
	OperationTypeID uint    `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func CreateTransaction(data *CreateTransactionData) error {
	DB := config.GetDB()
	logger := config.GetLogger("createTransactionService")

	accountRepository := repository.NewAccountRepository(DB)
	account := models.Account{}

	if err := accountRepository.FindByID(&account, fmt.Sprintf("%d", data.AccountID)); err != nil {
		logger.Errorf("account not found: %v", err)
		return fmt.Errorf("account does not exist")
	}

	operationTypeRepository := repository.NewOperationTypeRepository(DB)
	operationType := models.OperationType{}

	if err := operationTypeRepository.FindByID(&operationType, fmt.Sprintf("%d", data.OperationTypeID)); err != nil {
		logger.Errorf("operation type not found: %v", err)
		return fmt.Errorf("operation type does not exist")
	}

	if data.OperationTypeID == 4 {
		data.Amount = math.Abs(data.Amount)
	} else {
		data.Amount = -math.Abs(data.Amount)
	}

	transaction := models.Transaction{
		AccountID:       data.AccountID,
		OperationTypeID: data.OperationTypeID,
		Amount:          int(data.Amount * 100),
	}

	transactionRepository := repository.NewTransactionRepository(DB)

	if err := transactionRepository.Create(&transaction); err != nil {
		logger.Errorf("failed to create transaction: %v", err)
		return err
	}

	return nil
}
