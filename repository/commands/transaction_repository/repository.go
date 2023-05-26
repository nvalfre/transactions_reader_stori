package transaction_repository

import (
	"database/sql"
)

// DatabaseRepo handles interactions with the database
type TransactionDatabaseRepo struct {
	Db *sql.DB
}
