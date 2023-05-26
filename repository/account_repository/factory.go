package account_repository

import (
	"transactions_reader_stori/repository"
)

func NewAccountDatabaseRepo(db *repository.DatabaseRepo) AccountRepository {
	return &AccountDatabaseRepo{Db: db.Db}
}
