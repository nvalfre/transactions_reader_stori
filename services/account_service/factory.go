package account_service

import (
	"transactions_reader_stori/repository/account_repository"
)

// NewAccountService creates a new instance of AccountService
func NewAccountService(repo account_repository.AccountRepository) AccountServiceI {
	return &AccountService{
		repo: repo,
	}
}
