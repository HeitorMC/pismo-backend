package config

import (
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

func ConnectPostgres() (*gorm.DB, error) {
	logger := GetLogger("postgreSQL")
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		return nil, err
	}

	logger.Info("Database connection established successfully")

	return db, nil
}
