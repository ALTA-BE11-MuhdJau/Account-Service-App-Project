package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DBconn() *sql.DB {
	dbConn := os.Getenv("DB_CONN")
	db, err := sql.Open("mysql", dbConn)

	if err != nil {
		log.Fatal("Error sql Open.", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("Error connect to DB.", errPing.Error())
	} else {
		fmt.Println("Success connect to DB.")
	}
	return db
}
