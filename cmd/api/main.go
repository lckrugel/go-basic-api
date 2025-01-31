package main

import (
	"log"

	"github.com/lckrugel/go-stock/internal/api"
	"github.com/lckrugel/go-stock/internal/config"
)

func main() {
	cfg, cfgError := config.LoadConfig()
	if cfgError != nil {
		log.Fatal("Failed to load configuration: ", cfgError)
	}

	httpServer := api.NewHTTPServer(cfg.AppConfig)
	serverError := httpServer.Start()
	if serverError != nil {
		log.Fatal("Failed to start server: ", serverError)
	}
}
