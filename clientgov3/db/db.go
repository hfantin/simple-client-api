package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(user, password, dbname string) *sql.DB {
	dbUri := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	// db.SetConnMaxLifetime(1000)
	return db
}
