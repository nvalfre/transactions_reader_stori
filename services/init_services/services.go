package init_services

import (
	"transactions_reader_stori/repository/commands/account_repository"
	"transactions_reader_stori/repository/commands/transaction_repository"
	"transactions_reader_stori/services/account_service"
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/file_service"
	"transactions_reader_stori/services/file_service/file_service_content_reader"
	"transactions_reader_stori/services/file_service/file_service_summary_generator"
	"transactions_reader_stori/services/transaction_service"
)

type ServicesInitI interface {
	Init(
		transactionRepository transaction_repository.TransactionRepository,
		accountDatabaseRepo account_repository.AccountRepository,
	)
}
type Services struct {
	AccountService     account_service.AccountServiceI
	TransactionService transaction_service.TransactionServiceI
	FileService        file_service.FileServiceI
	EmailService       email_service.EmailServiceI
}

func InitWith(
	accountServiceFactory account_service.AccountServiceFactoryI,
	transactionServiceFactory transaction_service.TransactionServiceFactoryI,
	fileServiceFactory file_service.FileServiceFactoryI,
	emailServiceFactory email_service.EmailServiceFactoryI,
	transactionRepository transaction_repository.TransactionRepository,
	accountDatabaseRepo account_repository.AccountRepository,
) *Services {
	accountService := accountServiceFactory.NewAccountService(accountDatabaseRepo)
	transactionService := transactionServiceFactory.NewTransactionService(transactionRepository, accountService)
	emailService := emailServiceFactory.NewEmailServiceDefault()

	fileContentReaderUseCase := file_service_content_reader.NewFileContentReaderUseCase()
	fileSummaryGeneratorUseCase := file_service_summary_generator.NewFileSummaryGeneratorUseCase(transactionService, emailService)

	fileService := fileServiceFactory.NewFileService(fileContentReaderUseCase, fileSummaryGeneratorUseCase)

	return &Services{
		AccountService:     accountService,
		TransactionService: transactionService,
		FileService:        fileService,
		EmailService:       emailService,
	}
}
