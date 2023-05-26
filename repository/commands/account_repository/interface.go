package account_repository

import (
	"transactions_reader_stori/domain/dao"
)

type AccountRepository interface {
	SaveAccount(account *dao.Account) error
	UpdateAccountBalance(account *dao.Account) error
	GetAccountById(id string) (*dao.Account, error)
}
