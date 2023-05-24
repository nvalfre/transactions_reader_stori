package file_service_summary_generator

import (
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/transaction_service"
)

func NewFileSummaryGeneratorUseCase(transactionService transaction_service.TransactionServiceI, emailService email_service.EmailServiceI) FileSummaryGeneratorUseCaseI {
	return &FileSummaryGeneratorUseCase{
		transactionService: transactionService,
		emailService:       emailService,
	}
}
