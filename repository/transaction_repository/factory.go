package transaction_repository

import (
	"transactions_reader_stori/repository"
)

func NewTransactionDatabaseRepo(db *repository.DatabaseRepo) TransactionRepository {
	return &TransactionDatabaseRepo{Db: db.Db}
}
