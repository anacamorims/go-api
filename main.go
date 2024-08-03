package main

import (
	"github.com/anacamorims/go-api.git/config"
	"github.com/anacamorims/go-api.git/router"
)

var(
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//initialize Router
	router.Initialize()
}
