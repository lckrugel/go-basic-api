package application

import (
	"context"
	"errors"

	"github.com/lckrugel/go-basic-api/internal/api"
	"github.com/lckrugel/go-basic-api/internal/config"
	"github.com/lckrugel/go-basic-api/internal/database"
)

type Application struct {
	Db *database.DB
	// Logger *logger.Logger
	HttpClient *api.ServerHTTP
}

func NewApplication() (*Application, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		errMsg := "Failed to load config: " + err.Error()
		return nil, errors.New(errMsg)
	}

	server := api.NewHTTPServer(cfg.AppConfig)
	db, err := database.NewDBConnection(cfg.DatabaseConfig)
	if err != nil {
		errMsg := "Failed to start database: " + err.Error()
		return nil, errors.New(errMsg)
	}

	return &Application{
		Db:         db,
		HttpClient: server,
	}, nil
}

func (app *Application) Shutdown(ctx context.Context) error {
	if err := app.HttpClient.Close(ctx); err != nil {
		return err
	}

	if err := app.Db.Close(ctx); err != nil {
		return err
	}
	return nil
}
