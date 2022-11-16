package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connection() *sql.DB {
	var connectionString = os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("Error connect to db", errPing.Error())
	} else {
		fmt.Println("successful connection")
	}
	return db

}
