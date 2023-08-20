package app

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type DBConnection struct {
	Conn *pgx.Conn
}

func ConnectToDataBase(credentials string) *DBConnection {
	conn, err := pgx.Connect(context.Background(), credentials)
	if err != nil {
		Logger.Fatal(err.Error())
	}

	dbConnection := DBConnection{
		Conn: conn,
	}

	return &dbConnection
}
