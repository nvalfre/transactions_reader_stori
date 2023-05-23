package transaction_service

import "transactions_reader_stori/repository"

// TransactionService handles transaction_service-related operations
type TransactionService struct {
	repo *repository.DatabaseRepo
}
