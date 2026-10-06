package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(dataBaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dataBaseUrl)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	// read database pooling
	db.SetMaxIdleConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.ping: %w", err)
	}
	return db, nil
}
