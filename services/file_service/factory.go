package file_service

import (
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/transaction_service"
)

// NewFileService creates a new instance of FileService
func NewFileService(transactionService *transaction_service.TransactionService, emailService *email_service.EmailService) *FileService {
	return &FileService{
		transactionService: transactionService,
		emailService:       emailService,
	}
}
