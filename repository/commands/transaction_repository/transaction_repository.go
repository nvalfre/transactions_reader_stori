package transaction_repository

import (
	"database/sql"
	"log"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/domain/dao"
	"transactions_reader_stori/utils"
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
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
		return 0, err
	}
	return totalBalance, nil
}

// GetTransactionSummary retrieves the transaction summary from the database
func (r *TransactionDatabaseRepo) GetTransactionSummary(accountID string) ([]domain.TransactionSummary, error) {
	query := "SELECT id, account_id, date, amount, is_credit FROM TRANSACTIONS WHERE account_id = ?"

	rows, err := r.Db.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []dao.Transaction
	for rows.Next() {
		var transaction dao.Transaction
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Date, &transaction.Amount, &transaction.IsCredit)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return r.buildTransactionSummary(transactions), nil
}

func (r *TransactionDatabaseRepo) buildTransactionSummary(transaction []dao.Transaction) []domain.TransactionSummary {
	transactionSummary := make([]domain.TransactionSummary, 0)
	for _, t := range transaction {
		date, err := utils.ParseDateWithLayout("2006-01-02 15:04:05 -0700 MST", t.Date)
		if err != nil {
			log.Fatal(err)
		}
		ts := domain.TransactionSummary{
			ID:     int(t.ID),
			Month:  date,
			Amount: t.Amount,
		}
		transactionSummary = append(transactionSummary, ts)
	}
	return transactionSummary
}

// GetAverageCredit retrieves the average credit amount from the database
func (r *TransactionDatabaseRepo) GetAverageCredit() (float64, error) {
	var averageCredit float64
	query := "SELECT AVG(amount) FROM TRANSACTIONS WHERE is_credit = ?"
	err := r.Db.QueryRow(query, true).Scan(&averageCredit)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return 0, err
	}
	return averageDebit, nil
}
