package app

import (
	"database/sql"
)

type DBConnection struct {
	Db *sql.DB
}

func ConnectToDataBase(credentials string) *DBConnection {
	db, err := sql.Open("pgx", credentials)
	if err != nil {
		Logger.Fatal("error happened while opening connection to db")
	}

	dbConnection := DBConnection{
		Db: db,
	}

	return &dbConnection
}
