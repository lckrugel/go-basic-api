package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lckrugel/go-basic-api/internal/app/middleware"
	"github.com/lckrugel/go-basic-api/internal/container"
	"github.com/lckrugel/go-basic-api/internal/server"
)

func main() {
	appContainer, err := container.NewAppContainer()
	if err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	controllers := server.Controllers{
		UserController: appContainer.UserController,
	}

	// Initialize the application server
	router := server.NewRouter()
	router.RegisterRoutes(&controllers)
	globalMiddlewareChain := middleware.CreateChain(middleware.Logging)
	server := server.NewHTTPServer(appContainer.Config.AppConfig, globalMiddlewareChain(router.Mux))

	// Set up context to notify for OS interrupt signals.
	shutdownCtx, shutdownCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer shutdownCancel()

	// Start the server
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server encountered an error: %v", err)
		}
	}()

	// Wait for the shutdown signal
	<-shutdownCtx.Done()
	fmt.Println() // Print a newline to separate the shutdown message from the ^C
	log.Println("Shutting down...")

	// Shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}

	// Close the application container
	if err := appContainer.Close(ctx); err != nil {
		log.Printf("Failed to close application container: %v", err)
	}

	log.Printf("Shutdown complete")
}
