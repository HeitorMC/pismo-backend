package repository

import (
	"github.com/HeitorMC/pismo-backend/internal/models"
	"gorm.io/gorm"
)

type OperationTypeRepository interface {
	FindByID(operationType *models.OperationType, id string) error
}

type operationTypeRepository struct {
	db *gorm.DB
}

func NewOperationTypeRepository(db *gorm.DB) OperationTypeRepository {
	return &operationTypeRepository{db: db}
}

func (r *operationTypeRepository) FindByID(operationType *models.OperationType, id string) error {
	err := r.db.First(&operationType, id).Error
	if err != nil {
		return err
	}

	return nil
}
