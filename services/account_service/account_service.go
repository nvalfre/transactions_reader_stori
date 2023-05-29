package account_service

import "transactions_reader_stori/domain/dao"

// ProcessFileContent processes the file and saves the transactions to the database
func (s *AccountService) SaveAccount(acccount *dao.Account) error {
	err := s.repo.SaveAccount(acccount)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountService) UpdateAccountBalance(account *dao.Account) error {
	err := r.repo.UpdateAccountBalance(account)
	if err != nil {
		return err
	}
	return nil
}

// GetAccount generates the summary information
func (s *AccountService) GetAccount(id int) (*dao.Account, error) {
	account, err := s.repo.GetAccountById(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
