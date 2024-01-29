package main

import (
	"github.com/steniocg/gopportunities/config"
	"github.com/steniocg/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config inicialization error: %v", err)
		return
	}
	//Inicialize Router
	router.Initialize()

}
