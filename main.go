package main

import (
	"montravail/config"
	"montravail/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize error: %v", err)
		return
	}
	// Initialize the router
	router.Initialize()
}
