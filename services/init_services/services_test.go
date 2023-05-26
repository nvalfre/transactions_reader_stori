package init_services

import (
	"testing"
	"transactions_reader_stori/repository/account_repository"
	"transactions_reader_stori/repository/transaction_repository"
	"transactions_reader_stori/services/account_service"
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/file_service"
	"transactions_reader_stori/services/transaction_service"
)

func TestInitServices(t *testing.T) {
	// Create mock repositories
	transactionRepo := &transaction_repository.MockTransactionRepository{}
	accountRepo := &account_repository.MockAccountRepository{}

	// Initialize services
	s := Init(transactionRepo, accountRepo)

	// Check if services were initialized correctly
	if _, ok := s.AccountService.(*account_service.AccountService); !ok {
		t.Error("AccountService is not of the expected type")
	}
	if _, ok := s.TransactionService.(*transaction_service.TransactionService); !ok {
		t.Error("TransactionService is not of the expected type")
	}
	if _, ok := s.EmailService.(*email_service.EmailService); !ok {
		t.Error("EmailService is not of the expected type")
	}
	if _, ok := s.FileService.(*file_service.FileService); !ok {
		t.Error("FileService is not of the expected type")
	}
}
