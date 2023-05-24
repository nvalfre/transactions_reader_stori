package transaction_repository

import (
	"database/sql"
	"log"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/domain/dao"
)

// SaveTransaction saves a transaction_service to the database
func (r *TransactionDatabaseRepo) SaveTransaction(transaction *dao.Transaction) error {
	query := "INSERT INTO TRANSACTIONS (account_id, date, amount, is_credit) VALUES (?, ?, ?, ?)"
	result, err := r.Db.Exec(query, transaction.AccountID, transaction.Date, transaction.Amount, transaction.IsCredit)
	if err != nil {
		log.Fatal(err)
		return err
	}

	transactionID, _ := result.LastInsertId()
	transaction.ID = uint(int(transactionID))
	return nil
}

func (r *TransactionDatabaseRepo) UpdateTransaction(transaction *dao.Transaction) error {
	query := "UPDATE TRANSACTIONS SET date = ?, amount = ?, is_credit = ? WHERE id = ?"
	_, err := r.Db.Exec(query, transaction.Date, transaction.Amount, transaction.IsCredit, transaction.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *TransactionDatabaseRepo) GetTransactionByDateAndAccountID(date string, accountID uint) (*dao.Transaction, error) {
	query := "SELECT id, date, amount, is_credit, account_id FROM TRANSACTIONS WHERE date = ? AND account_id = ?"
	row := r.Db.QueryRow(query, date, accountID)

	var transaction dao.Transaction
	err := row.Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.IsCredit, &transaction.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &transaction, nil
}

// GetTotalBalance retrieves the total balance from the database
func (r *TransactionDatabaseRepo) GetTotalBalance() (float64, error) {
	var totalBalance float64
	query := "SELECT SUM(amount) FROM TRANSACTIONS"
	err := r.Db.QueryRow(query).Scan(&totalBalance)
	if err != nil {
		return 0, err
	}
	return totalBalance, nil
}

// GetTransactionSummary retrieves the transaction summary from the database
func (r *TransactionDatabaseRepo) GetTransactionSummary() ([]domain.TransactionSummary, error) {
	query := "SELECT MONTH(date) AS month, COUNT(*) AS num_of_trans FROM TRANSACTIONS GROUP BY month ORDER BY month"
	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactionSummary []domain.TransactionSummary
	for rows.Next() {
		var summary domain.TransactionSummary
		err := rows.Scan(&summary.Month, &summary.NumOfTrans)
		if err != nil {
			return nil, err
		}
		transactionSummary = append(transactionSummary, summary)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactionSummary, nil
}

// GetAverageCredit retrieves the average credit amount from the database
func (r *TransactionDatabaseRepo) GetAverageCredit() (float64, error) {
	var averageCredit float64
	query := "SELECT AVG(amount) FROM TRANSACTIONS WHERE is_credit = ?"
	err := r.Db.QueryRow(query, true).Scan(&averageCredit)
	if err != nil {
		return 0, err
	}
	return averageCredit, nil
}

// GetAverageDebit retrieves the average debit amount from the database
func (r *TransactionDatabaseRepo) GetAverageDebit() (float64, error) {
	var averageDebit float64
	query := "SELECT AVG(amount) FROM TRANSACTIONS WHERE is_credit = ?"
	err := r.Db.QueryRow(query, false).Scan(&averageDebit)
	if err != nil {
		return 0, err
	}
	return averageDebit, nil
}
