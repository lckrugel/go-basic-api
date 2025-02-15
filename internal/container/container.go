package container

import (
	"context"
	"errors"

	"github.com/lckrugel/go-basic-api/internal/app/controller"
	"github.com/lckrugel/go-basic-api/internal/app/service"
	"github.com/lckrugel/go-basic-api/internal/config"
	"github.com/lckrugel/go-basic-api/internal/database"
	"github.com/lckrugel/go-basic-api/internal/repository/sqlc"
)

type AppContainer struct {
	// Logger *logger.Logger
	db *database.DB

	Config         *config.Config
	UserService    *service.UserService
	UserController *controller.UserController
}

func NewAppContainer() (*AppContainer, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		errMsg := "Failed to load config: " + err.Error()
		return nil, errors.New(errMsg)
	}

	db, err := database.NewDBConnection(cfg.DatabaseConfig)
	if err != nil {
		errMsg := "Failed to start database: " + err.Error()
		return nil, errors.New(errMsg)
	}

	// logger := logger.NewLogger(cfg.LoggerConfig)

	queries := sqlc.New(db)
	userRepo := sqlc.NewSQLCUserRepository(queries)
	userSvc := service.NewUserService(userRepo)
	userCtl := controller.NewUserController(userSvc)

	return &AppContainer{
		// Logger: logger,
		db:             db,
		Config:         cfg,
		UserService:    userSvc,
		UserController: userCtl,
	}, nil
}

func (c *AppContainer) Close(ctx context.Context) error {
	return c.db.Close(ctx)
}
