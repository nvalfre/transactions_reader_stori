package account_repository

import (
	"database/sql"
)

// DatabaseRepo handles interactions with the database
type AccountDatabaseRepo struct {
	Db *sql.DB
}
