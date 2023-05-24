package repository

import (
	"database/sql"
	"transactions_reader_stori/domain/dao"
)

type AccountRepository interface {
	SaveAccount(account *dao.Account) error
	GetAccountById(id string) (*dao.Account, error)
}

type TransactionRepository interface {
	SaveTransaction(transaction *dao.Transaction) error
}

// DatabaseRepo handles interactions with the database
type DatabaseRepo struct {
	Db *sql.DB
}
