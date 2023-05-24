package transaction_service

import (
	"transactions_reader_stori/repository"
	"transactions_reader_stori/services/account_service"
)

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(repo *repository.DatabaseRepo, accountService account_service.AccountServiceI) TransactionServiceI {
	return &TransactionService{
		accountService: accountService,
		repo:           repo,
	}
}
