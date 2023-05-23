package repository

import (
	"transactions_reader_stori/domain/dao"
)

// SaveTransaction saves a transaction_service to the database
func (r *DatabaseRepo) SaveTransaction(transaction *dao.Transaction) error {
	return r.Db.Create(transaction).Error
}
