package repository

import (
	"log"
	"transactions_reader_stori/domain/dao"
)

// SaveAccount saves an account to the database
func (r *DatabaseRepo) SaveAccount(account *dao.Account) error {
	// Insert the account into the database
	result, err := r.Db.Exec("INSERT INTO accounts (name, email, balance) VALUES (?, ?)", account.Name, account.Email, account.Balance)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Get the inserted account ID
	accountID, _ := result.LastInsertId()
	account.ID = uint(int(accountID))
	return nil
}
