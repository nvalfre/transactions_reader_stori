package transaction_service

import (
	"transactions_reader_stori/repository"
	"transactions_reader_stori/services/account_service"
)

// TransactionService handles transaction_service-related operations
type TransactionService struct {
	accountService account_service.AccountServiceI
	repo           *repository.DatabaseRepo
}
