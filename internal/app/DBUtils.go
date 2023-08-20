package app

import (
	"database/sql"
	"fmt"
)

type DBCredentials struct {
	user     string
	password string
	host     string
	dbname   string
}

func ConnectToDataBase(credentials *DBCredentials) {
	ps := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		credentials.host, credentials.user, credentials.password, credentials.dbname)

	db, err := sql.Open("pgx", ps)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			Logger.Fatal("error happened while closing db connection")
		}
	}(db)
}
