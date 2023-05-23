package file_service

import (
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/transaction_service"
)

// FileService handles file processing HTTP endpoint
type FileService struct {
	transactionService *transaction_service.TransactionService
	emailService       *email_service.EmailService
}
