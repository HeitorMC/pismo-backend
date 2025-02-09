package main

import (
	"github.com/HeitorMC/pismo-backend/config"
	"github.com/HeitorMC/pismo-backend/internal/models"
)

var (
	logger *config.Logger
)

func init() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
}

func main() {
	logger.Info("Starting migration...")

	db := config.GetDB()

	err := db.AutoMigrate(&models.Account{}, &models.OperationType{}, &models.Transaction{})

	if err != nil {
		logger.Errorf("Migration failed: %v", err)
		return
	}

	logger.Info("Migration completed!")

	logger.Info("Seeding operation types...")

	models.SeedOperationTypes(db)

	logger.Info("Operation types seeded!")
}
