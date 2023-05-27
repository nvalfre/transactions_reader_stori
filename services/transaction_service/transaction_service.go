package transaction_service

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/domain/dao"
	"transactions_reader_stori/services/builders"
	"transactions_reader_stori/utils"
	"transactions_reader_stori/validators"
)

// ProcessFileContent processes the file and saves the transactions to the database
func (s *TransactionService) ProcessFileContent(fileContent []byte, accountId string, accountName string) error {
	lines := strings.Split(string(fileContent), "\n")

	account, err := s.accountService.GetAccount(accountId)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if account.ID == 0 {
		newAccount := &dao.Account{Balance: 0, Name: accountName}
		if err := s.accountService.SaveAccount(newAccount); err != nil {
			return err
		}
		account = newAccount
	}

	for _, line := range lines[1:] {
		fields := strings.Split(line, ",")
		if len(fields) != 3 {
			continue
		}

		date := strings.TrimSpace(fields[1])
		amount := strings.TrimSpace(fields[2])

		transaction := dao.Transaction{
			Date:      date,
			Amount:    utils.ParseAmount(amount),
			IsCredit:  validators.IsCredit(amount),
			AccountID: account.ID,
		}

		existingTransaction, err := s.repo.GetTransactionByDateAndAccountID(date, account.ID)
		if err != nil {
			return err
		}

		if existingTransaction != nil {
			log.Printf("Transaction already exists for date: %s and account ID: %v. Performing update.", date, account.ID)

			transaction.ID = existingTransaction.ID
			if err := s.repo.UpdateTransaction(&transaction); err != nil {
				return err
			}
			if transaction.IsCredit {
				account.Balance += transaction.Amount
			} else {
				account.Balance -= transaction.Amount
			}

			if err := s.accountService.UpdateAccountBalance(account); err != nil {
				return err
			}
		} else {
			if transaction.IsCredit {
				account.Balance += transaction.Amount
			} else {
				account.Balance -= transaction.Amount
			}

			if err := s.accountService.SaveAccount(account); err != nil {
				return err
			}

			if err := s.repo.SaveTransaction(&transaction); err != nil {
				return err
			}
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

	transactionMetadata, err := s.getTransactionMetadata(err)
	if err != nil {
		return nil, err
	}

	summary := builders.BuildSummary(transactionMetadata)

	return summary, nil
}
func (s *TransactionService) getTransactionMetadata(err error) (*domain.TransactionMetadata, error) {
	totalBalance, err := s.repo.GetTotalBalance()
	if err != nil {
		return nil, err
	}

	transactionSummary, err := s.repo.GetTransactionSummary()
	if err != nil {
		return nil, err
	}

	averageCredit, err := s.repo.GetAverageCredit()
	if err != nil {
		return nil, err
	}

	averageDebit, err := s.repo.GetAverageDebit()
	if err != nil {
		return nil, err
	}
	return builders.BuildTransactionMetadata(totalBalance, transactionSummary, averageCredit, averageDebit), nil
}
