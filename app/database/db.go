package database

import (
	"fmt"
	"context"
	
	"github.com/jackc/pgx/v5"
	"cappuchinodb.com/main/app/config"
)

var DB *pgx.Conn

func ConnectDB() (*pgx.Conn, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/cappuchinodb",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
	)
	
	var err error
	DB, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return DB, nil
}
