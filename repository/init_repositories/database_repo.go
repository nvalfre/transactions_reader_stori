package init_repositories

import "database/sql"

type DatabaseRepo struct {
	Db *sql.DB
}
