package repository

import (
	"fmt"

	"github.com/HeitorMC/pismo-backend/internal/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *models.Account) error
	Find(account *models.Account, id string) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *models.Account) error {
	err := r.db.Create(account).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return fmt.Errorf("document number already being used")
		}
		return err
	}

	return nil
}

func (r *accountRepository) Find(account *models.Account, id string) error {
	err := r.db.First(account, id).Error
	if err != nil {
		return err
	}

	return nil
}
