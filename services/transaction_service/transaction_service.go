package transaction_service

import (
	"database/sql"
	"errors"
	"strings"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/domain/dao"
	"transactions_reader_stori/utils"
	"transactions_reader_stori/validators"
)

// ProcessFile processes the file and saves the transactions to the database
func (s *TransactionService) ProcessFile(fileContent []byte, accountId string, accountName string) error {
	lines := strings.Split(string(fileContent), "\n")

	account, err := s.accountService.GetAccount(accountId)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	for _, line := range lines[1:] {
		fields := strings.Split(line, ",")
		if len(fields) != 3 {
			continue
		}

		date := strings.TrimSpace(fields[1])
		amount := strings.TrimSpace(fields[2])

		var transaction dao.Transaction
		transaction.Date = date
		transaction.Amount = utils.ParseAmount(amount)
		transaction.IsCredit = validators.IsCredit(amount)

		if account.ID == 0 {
			newAccount := &dao.Account{Balance: transaction.Amount, Name: accountName}
			if err := s.repo.SaveAccount(newAccount); err != nil {
				return err
			}
			transaction.AccountID = newAccount.ID
			account = newAccount
		} else {
			transaction.AccountID = account.ID
			if transaction.IsCredit {
				account.Balance += transaction.Amount
			} else {
				account.Balance -= transaction.Amount
			}
			if err := s.repo.SaveAccount(account); err != nil {
				return err
			}
		}

		if err := s.repo.SaveTransaction(&transaction); err != nil {
			return err
		}
	}

	return nil
}

// GenerateSummary generates the summary information
func (s *TransactionService) GenerateSummary(accountId string) (*domain.SummaryVO, error) {
	account, err := s.accountService.GetAccount(accountId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	if account.ID == 0 {
		return nil, errors.New("invalid_account")
	}

	var totalBalance float64
	if err := s.repo.Db.QueryRow("SELECT SUM(amount) FROM TRANSACTIONS").Scan(&totalBalance); err != nil {
		return nil, err
	}

	rows, err := s.repo.Db.Query("SELECT MONTH(date) AS month, COUNT(*) AS num_of_trans FROM TRANSACTIONS GROUP BY month ORDER BY month")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactionSummary := make(map[string]int)
	for rows.Next() {
		var month string
		var numOfTrans int
		err := rows.Scan(&month, &numOfTrans)
		if err != nil {
			return nil, err
		}
		transactionSummary[month] = numOfTrans
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var averageCredit, averageDebit float64
	if err := s.repo.Db.QueryRow("SELECT AVG(amount) FROM TRANSACTIONS WHERE is_credit = ?", true).Scan(&averageCredit); err != nil {
		return nil, err
	}
	if err := s.repo.Db.QueryRow("SELECT AVG(amount) FROM TRANSACTIONS WHERE is_credit = ?", false).Scan(&averageDebit); err != nil {
		return nil, err
	}

	summary := &domain.SummaryVO{
		TotalBalance:       totalBalance,
		TransactionSummary: transactionSummary,
		AverageCredit:      averageCredit,
		AverageDebit:       averageDebit,
	}

	return summary, nil
}
