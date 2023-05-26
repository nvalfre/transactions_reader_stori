package repository

import "database/sql"

type DatabaseRepo struct {
	Db *sql.DB
}
