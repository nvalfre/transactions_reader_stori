package account_repository

import (
	"gorm.io/gorm"
	"transactions_reader_stori/domain/dao"
)

type MockAccountRepository struct {
}

func (mockAccountRepository MockAccountRepository) name() {

}
func (mockAccountRepository MockAccountRepository) SaveAccount(account *dao.Account) error {
	return nil
}
func (mockAccountRepository MockAccountRepository) UpdateAccountBalance(account *dao.Account) error {
	return nil
}
func (mockAccountRepository MockAccountRepository) GetAccountById(id string) (*dao.Account, error) {
	return &dao.Account{
		Model:   gorm.Model{},
		Name:    "",
		Email:   "",
		Balance: 0,
	}, nil
}
