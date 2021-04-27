package models

import (
	"context"
	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

func InitDB(databaseUrl string) (*pgx.Conn, error) {
	var err error
	conn, _ = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return conn, err
	}
	return conn, nil
}

func GetDB() *pgx.Conn {
	return conn
}

func CloseDB() error {
	return conn.Close(context.Background())
}
