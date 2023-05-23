package transaction_service

import (
	"gorm.io/gorm"
	"strings"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/domain/dao"
	"transactions_reader_stori/utils"
	"transactions_reader_stori/validators"
)

// ProcessFile processes the file and saves the transactions to the database
func (s *TransactionService) ProcessFile(fileContent []byte) error {
	lines := strings.Split(string(fileContent), "\n")

	var account dao.Account
	if err := s.repo.Db.First(&account).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
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
			newAccount := dao.Account{Balance: transaction.Amount}
			if err := s.repo.SaveAccount(&newAccount); err != nil {
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
			if err := s.repo.SaveAccount(&account); err != nil {
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
func (s *TransactionService) GenerateSummary() (*domain.Summary, error) {
	var account dao.Account
	if err := s.repo.Db.First(&account).Error; err != nil {
		return nil, err
	}

	var totalBalance float64
	if err := s.repo.Db.Model(&dao.Transaction{}).Select("SUM(amount)").Scan(&totalBalance).Error; err != nil {
		return nil, err
	}

	var transactionSummary []struct {
		Month      string
		NumOfTrans int64
	}
	if err := s.repo.Db.Select("strftime('%m', date) AS month, COUNT(*) AS num_of_trans").
		Group("month").
		Order("month").
		Find(&transactionSummary).Error; err != nil {
		return nil, err
	}

	var averageCredit, averageDebit float64
	if err := s.repo.Db.Model(&dao.Transaction{}).Select("AVG(amount)").
		Where("is_credit = ?", true).
		Scan(&averageCredit).Error; err != nil {
		return nil, err
	}
	if err := s.repo.Db.Model(&dao.Transaction{}).Select("AVG(amount)").
		Where("is_credit = ?", false).
		Scan(&averageDebit).Error; err != nil {
		return nil, err
	}

	summary := &domain.Summary{
		TotalBalance:       totalBalance,
		TransactionSummary: make(map[string]int),
		AverageCredit:      averageCredit,
		AverageDebit:       averageDebit,
	}

	for _, t := range transactionSummary {
		summary.TransactionSummary[t.Month] = int(t.NumOfTrans)
	}

	return summary, nil
}
