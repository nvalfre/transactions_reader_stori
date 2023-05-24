package repository

import (
	"log"
	"transactions_reader_stori/domain/dao"
)

// SaveAccount saves an account to the database
func (r *DatabaseRepo) SaveAccount(account *dao.Account) error {
	// Insert the account into the database
	result, err := r.Db.Exec("INSERT INTO ACCOUNTS (name, email, balance) VALUES (?, ?)", account.Name, account.Email, account.Balance)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Get the inserted account ID
	accountID, _ := result.LastInsertId()
	account.ID = uint(int(accountID))
	return nil
}

// GetAccountById retrieves an account from the database
func (r *DatabaseRepo) GetAccountById(id string) (*dao.Account, error) {
	account := &dao.Account{}

	row := r.Db.QueryRow("SELECT id, name, balance FROM ACCOUNTS WHERE id = ? LIMIT 1", id)
	err := row.Scan(&account.ID, &account.Name, &account.Balance)
	if err != nil {
		return nil, err
	}

	return account, nil
}
