package account_service

import "transactions_reader_stori/domain/dao"

type AccountServiceI interface {
	UpdateAccountBalance(acccount *dao.Account) error
	SaveAccount(acccount *dao.Account) error
	GetAccount(id int) (*dao.Account, error)
}
