package service

import (
	"github.com/HeitorMC/pismo-backend/config"
	"github.com/HeitorMC/pismo-backend/internal/models"
	"github.com/HeitorMC/pismo-backend/internal/repository"
)

func CreateAccount(account *models.Account) error {
	DB := config.GetDB()
	logger := config.GetLogger("createAccountService")

	accountRepository := repository.NewAccountRepository(DB)

	if err := accountRepository.Create(account); err != nil {
		logger.Errorf("failed to create account: %v", err)
		return err
	}

	return nil
}
