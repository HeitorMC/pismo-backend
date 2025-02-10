package main

import (
	"github.com/HeitorMC/pismo-backend/config"
	"github.com/HeitorMC/pismo-backend/internal/models"
)

func main() {
	logger := config.GetLogger("migration")

	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	logger.Info("Starting migration...")

	db := config.GetDB()

	err = db.AutoMigrate(&models.Account{}, &models.OperationType{}, &models.Transaction{})

	if err != nil {
		logger.Errorf("Migration failed: %v", err)
		return
	}

	logger.Info("Migration completed!")

	logger.Info("Seeding operation types...")

	models.SeedOperationTypes(db)

	logger.Info("Operation types seeded!")
}
