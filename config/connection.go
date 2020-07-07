package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, _ := sql.Open("mysql", "root:admin@tcp(localhost:3306)/digitalmeeting")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
