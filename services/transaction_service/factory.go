package transaction_service

import (
	"transactions_reader_stori/repository/commands/transaction_repository"
	"transactions_reader_stori/services/account_service"
)

type TransactionServiceFactoryI interface {
	NewTransactionService(repo transaction_repository.TransactionRepository, accountService account_service.AccountServiceI) TransactionServiceI
}

type TransactionServiceFactory struct{}

// NewTransactionService creates a new instance of TransactionService
func (factory TransactionServiceFactory) NewTransactionService(repo transaction_repository.TransactionRepository, accountService account_service.AccountServiceI) TransactionServiceI {
	return &TransactionService{
		accountService: accountService,
		repo:           repo,
	}
}
