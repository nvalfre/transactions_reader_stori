package repository

import (
	"database/sql"
	"log"
)

// DatabaseRepo handles interactions with the database
type DatabaseRepo struct {
	Db *sql.DB
}

// NewDatabaseRepo creates a new instance of DatabaseRepo
func NewDatabaseRepo(db *sql.DB) *DatabaseRepo {
	db, err := sql.Open("sqlite3", "transactions.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);

	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		account_id INTEGER,
		date TEXT,
		amount REAL,
		is_credit INTEGER,
		FOREIGN KEY(account_id) REFERENCES accounts(id)
	);
`)
	if err != nil {
		log.Fatal(err)
	}
	return &DatabaseRepo{
		Db: db,
	}
}
