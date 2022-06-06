package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "practice6"
)

var client *sql.DB

func OpenDBConnection() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
