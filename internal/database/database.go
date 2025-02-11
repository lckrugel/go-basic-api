package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/lckrugel/go-basic-api/internal/config"
)

type DB struct {
	conn *pgx.Conn
}

func NewDBConnection(cfg config.DatabaseConfig) (*DB, error) {
	ctx := context.Background()

	connStr := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		conn: conn,
	}, nil
}

func (db *DB) Close(ctx context.Context) error {
	return db.conn.Close(ctx)
}
