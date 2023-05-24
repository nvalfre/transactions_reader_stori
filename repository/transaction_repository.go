package repository

import (
	"log"
	"transactions_reader_stori/domain/dao"
)

// SaveTransaction saves a transaction_service to the database
func (r *DatabaseRepo) SaveTransaction(transaction *dao.Transaction) error {
	result, err := r.Db.Exec("INSERT INTO TRANSACTIONS (account_id, date, amount, is_credit) VALUES (?, ?, ?, ?)", transaction.AccountID, transaction.Date, transaction.Amount, transaction.IsCredit)
	if err != nil {
		log.Fatal(err)
		return err
	}

	transactionID, _ := result.LastInsertId()
	transaction.ID = uint(int(transactionID))
	return nil
}
