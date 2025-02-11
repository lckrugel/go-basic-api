package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lckrugel/go-basic-api/internal/application"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	go func() {
		if err := app.HttpClient.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server encountered an error: %v", err)
		}
	}()

	// Set up channel to listen for OS interrupt signals.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Create a context with timeout for the shutdown process.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatalf("Application shutdown failed: %v", err)
	}

	log.Println("Application stopped gracefully.")
}
