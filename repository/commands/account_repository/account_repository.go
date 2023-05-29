package account_repository

import (
	"log"
	"transactions_reader_stori/domain/dao"
)

// SaveAccount saves an account to the database
func (r *AccountDatabaseRepo) SaveAccount(account *dao.Account) error {
	// Insert the account into the database
	query := "INSERT INTO ACCOUNTS (name, email, balance) VALUES (?, ?, ?)"

	result, err := r.Db.Exec(query, account.Name, account.Email, account.Balance)
	if err != nil {
		log.Println(err)
		return err
	}

	// Get the inserted account ID
	accountID, _ := result.LastInsertId()
	account.ID = uint(int(accountID))
	return nil
}

func (r *AccountDatabaseRepo) UpdateAccountBalance(account *dao.Account) error {
	query := "UPDATE ACCOUNTS SET balance = ? WHERE id = ?"

	_, err := r.Db.Exec(query, account.Balance, account.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetAccountById retrieves an account from the database
func (r *AccountDatabaseRepo) GetAccountById(id int) (*dao.Account, error) {
	account := &dao.Account{}
	query := "SELECT id, name, balance FROM ACCOUNTS WHERE id = ? LIMIT 1"

	row := r.Db.QueryRow(query, id)
	err := row.Scan(&account.ID, &account.Name, &account.Balance)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return account, nil
}
