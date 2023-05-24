package account_service

import "transactions_reader_stori/domain/dao"

type AccountServiceI interface {
	SaveAccount(fileContent []byte) error
	GetAccount(id string) (*dao.Account, error)
}
