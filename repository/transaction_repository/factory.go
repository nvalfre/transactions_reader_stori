package transaction_repository

import (
	"transactions_reader_stori/repository/init_repositories"
)

func NewTransactionDatabaseRepo(db *init_repositories.DatabaseRepo) TransactionRepository {
	return &TransactionDatabaseRepo{Db: db.Db}
}
