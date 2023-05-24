package services

import (
	"transactions_reader_stori/repository"
	"transactions_reader_stori/services/account_service"
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/file_service"
	"transactions_reader_stori/services/file_service/file_service_content_reader"
	"transactions_reader_stori/services/file_service/file_service_summary_generator"
	"transactions_reader_stori/services/transaction_service"
)

type Services struct {
	AccountService     account_service.AccountServiceI
	TransactionService transaction_service.TransactionServiceI
	FileService        file_service.FileServiceI
	EmailService       email_service.EmailServiceI
}

func InitServices(db *repository.DatabaseRepo) *Services {
	accountService := account_service.NewAccountService(db)
	transactionService := transaction_service.NewTransactionService(db, accountService)
	emailService := email_service.NewEmailServiceDefault()

	fileContentReaderUseCase := file_service_content_reader.NewFileContentReaderUseCase()
	fileSummaryGeneratorUseCase := file_service_summary_generator.NewFileSummaryGeneratorUseCase(transactionService, emailService)

	fileService := file_service.NewFileService(fileContentReaderUseCase, fileSummaryGeneratorUseCase)

	return &Services{
		AccountService:     accountService,
		TransactionService: transactionService,
		FileService:        fileService,
		EmailService:       emailService,
	}
}
