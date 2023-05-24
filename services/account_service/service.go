package account_service

import "transactions_reader_stori/repository"

// TransactionService handles transaction_service-related operations
type AccountService struct {
	repo *repository.DatabaseRepo
}
