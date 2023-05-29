package main

import (
	"database/sql"
	"log"
	"time"
	"transactions_reader_stori/app_config/initializer"
)

func main() {
	appComponentsInitializer := initializer.AppComponentsInitializer{}

	db := openDbConnection()
	defer db.Close()
	appComponentsInitializer.Init(db)
}

func openDbConnection() *sql.DB {
	db, err := sql.Open("sqlite", "transactions.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
