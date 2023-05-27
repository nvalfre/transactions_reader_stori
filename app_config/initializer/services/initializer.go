package services

import (
	"transactions_reader_stori/services/account_service"
	"transactions_reader_stori/services/email_service"
	"transactions_reader_stori/services/file_service"
	"transactions_reader_stori/services/transaction_service"
)

type AppServicesComponentsInitializerI interface {
	InitServicesFactories() (*account_service.AccountServiceFactory, *transaction_service.TransactionServiceFactory, *file_service.FileServiceFactory, *email_service.EmailServiceFactory)
}

type AppServicesComponentsInitializer struct {
}

func (initializer AppServicesComponentsInitializer) InitServicesFactories() (*account_service.AccountServiceFactory, *transaction_service.TransactionServiceFactory, *file_service.FileServiceFactory, *email_service.EmailServiceFactory) {
	accountServiceFactory := &account_service.AccountServiceFactory{}
	transactionServiceFactory := &transaction_service.TransactionServiceFactory{}
	fileServiceFactory := &file_service.FileServiceFactory{}
	emailServiceFactory := &email_service.EmailServiceFactory{}
	return accountServiceFactory, transactionServiceFactory, fileServiceFactory, emailServiceFactory
}
