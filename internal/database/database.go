package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/lckrugel/go-basic-api/internal/config"
)

type DB struct {
	*pgx.Conn
}

func NewDBConnection(cfg config.DatabaseConfig) (*DB, error) {
	ctx := context.Background()

	connStr := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}
