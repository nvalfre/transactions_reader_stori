package account_service

import "transactions_reader_stori/domain/dao"

// ProcessFile processes the file and saves the transactions to the database
func (s *AccountService) SaveAccount(fileContent []byte) error {
	return nil
}

// GetAccount generates the summary information
func (s *AccountService) GetAccount(id string) (*dao.Account, error) {
	account, err := s.repo.GetAccountById(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
