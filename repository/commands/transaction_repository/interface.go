package transaction_repository

import (
	"transactions_reader_stori/domain"
	"transactions_reader_stori/domain/dao"
)

type TransactionRepository interface {
	SaveTransaction(transaction *dao.Transaction) error
	UpdateTransaction(transaction *dao.Transaction) error
	GetTransactionByDateAndAccountID(date string, accountID uint) (*dao.Transaction, error)
	GetTotalBalance() (float64, error)
	GetTransactionSummary(accountID string) ([]domain.TransactionSummary, error)
	GetAverageCredit() (float64, error)
	GetAverageDebit() (float64, error)
}
