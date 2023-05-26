package account_service

import (
	"transactions_reader_stori/repository/account_repository"
)

// TransactionService handles transaction_service-related operations
type AccountService struct {
	repo account_repository.AccountRepository
}
