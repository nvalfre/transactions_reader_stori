package account_repository

import (
	"transactions_reader_stori/repository/init_repositories"
)

func NewAccountDatabaseRepo(db *init_repositories.DatabaseRepo) AccountRepository {
	return &AccountDatabaseRepo{Db: db.Db}
}
