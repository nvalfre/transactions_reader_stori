package repository

import (
	"transactions_reader_stori/domain/dao"
)

// SaveAccount saves an account to the database
func (r *DatabaseRepo) SaveAccount(account *dao.Account) error {
	return r.Db.Create(account).Error
}
