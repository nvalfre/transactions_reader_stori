package transaction_repository_factory

import (
	transaction_repository2 "transactions_reader_stori/repository/commands/transaction_repository"
	"transactions_reader_stori/repository/init_repositories"
)

func NewTransactionDatabaseRepo(db *init_repositories.DatabaseRepo) transaction_repository2.TransactionRepository {
	return &transaction_repository2.TransactionDatabaseRepo{Db: db.Db}
}
