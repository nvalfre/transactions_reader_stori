package account_service

import (
	"transactions_reader_stori/repository"
)

// NewAccountService creates a new instance of AccountService
func NewAccountService(repo *repository.DatabaseRepo) AccountServiceI {
	return &AccountService{
		repo: repo,
	}
}
