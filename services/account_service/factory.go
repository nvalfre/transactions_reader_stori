package account_service

import (
	"transactions_reader_stori/repository/commands/account_repository"
)

type AccountServiceFactoryI interface {
	NewAccountService(repo account_repository.AccountRepository) AccountServiceI
}

type AccountServiceFactory struct{}

// NewAccountService creates a new instance of AccountService
func (factory *AccountServiceFactory) NewAccountService(repo account_repository.AccountRepository) AccountServiceI {
	return &AccountService{
		repo: repo,
	}
}
