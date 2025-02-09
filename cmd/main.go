package main

import (
	"github.com/HeitorMC/pismo-backend/config"
	"github.com/HeitorMC/pismo-backend/internal/router"
)

func main() {

	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
