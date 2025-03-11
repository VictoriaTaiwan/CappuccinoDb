package data

import (
	"context"
	"github.com/jackc/pgx/v5"
)

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres" 
	Password = "cappuchino_admin_fhgr769365"
	Dbname   = "cappuchinodb"
)

func ConnectDB() (*pgx.Conn, error) {
	address := "postgres://" + User + ":" + Password + "@" + Host + ":" + Port + "/cappuchinodb"
	conn, err := pgx.Connect(context.Background(), address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

