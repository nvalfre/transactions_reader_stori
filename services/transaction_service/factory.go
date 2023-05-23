package transaction_service

import (
	"transactions_reader_stori/repository"
)

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(repo *repository.DatabaseRepo) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}
