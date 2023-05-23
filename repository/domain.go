package repository

import (
	"gorm.io/gorm"
	"transactions_reader_stori/domain/dao"
)

// DatabaseRepo handles interactions with the database
type DatabaseRepo struct {
	Db *gorm.DB
}

// NewDatabaseRepo creates a new instance of DatabaseRepo
func NewDatabaseRepo(db *gorm.DB) *DatabaseRepo {
	return &DatabaseRepo{
		Db: db,
	}
}

// SaveAccount saves an account to the database
func (r *DatabaseRepo) SaveAccount(account *dao.Account) error {
	return r.Db.Create(account).Error
}

// SaveTransaction saves a transaction_service to the database
func (r *DatabaseRepo) SaveTransaction(transaction *dao.Transaction) error {
	return r.Db.Create(transaction).Error
}
