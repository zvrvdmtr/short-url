package models

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type DBConnect interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

var conn DBConnect

func InitDB(databaseUrl string) (DBConnect, error) {
	var err error
	conn, _ = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return conn, err
	}
	return conn, nil
}
