package service

import (
	"github.com/HeitorMC/pismo-backend/config"
	"github.com/HeitorMC/pismo-backend/internal/models"
	"github.com/HeitorMC/pismo-backend/internal/repository"
)

func FindAccount(account *models.Account, id string) error {
	DB := config.GetDB()
	logger := config.GetLogger("findAccountService")

	accountRepository := repository.NewAccountRepository(DB)

	if err := accountRepository.FindByID(account, id); err != nil {
		logger.Errorf("failed to find account: %v", err)
		return err
	}

	return nil

}
