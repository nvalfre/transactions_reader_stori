package init_repositories

import (
	"database/sql"
	"transactions_reader_stori/repository/commands/account_repository"
	"transactions_reader_stori/repository/commands/transaction_repository"
)

type DatabaseRepo struct {
	Db *sql.DB
}

type DatabaseRepoCommands struct {
	TransactionRepository transaction_repository.TransactionRepository
	AccountRepository     account_repository.AccountRepository
}
