package repository

import (
	"transactions_reader_stori/domain/dao"
)

type Repository interface {
	SaveAccount(account *dao.Account) error
	SaveTransaction(transaction *dao.Transaction) error
}
