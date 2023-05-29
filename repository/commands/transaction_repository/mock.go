package transaction_repository

import (
	"time"
	"transactions_reader_stori/domain"
	"transactions_reader_stori/domain/dao"
)

type MockTransactionRepository struct {
}

func (mockTransactionRepository MockTransactionRepository) name() {

}
func (mockTransactionRepository MockTransactionRepository) SaveTransaction(transaction *dao.Transaction) error {
	return nil
}
func (mockTransactionRepository MockTransactionRepository) UpdateTransaction(transaction *dao.Transaction) error {
	return nil
}
func (mockTransactionRepository MockTransactionRepository) GetTransactionByDateAndAccountID(date string, accountID uint) (*dao.Transaction, error) {
	return &dao.Transaction{}, nil
}
func (mockTransactionRepository MockTransactionRepository) GetTotalBalance() (float64, error) {
	return 1, nil
}
func (mockTransactionRepository MockTransactionRepository) GetTransactionSummary() ([]domain.TransactionSummary, error) {
	return []domain.TransactionSummary{{
		Month:  time.Now(),
		Amount: 1,
	}}, nil
}
func (mockTransactionRepository MockTransactionRepository) GetAverageCredit() (float64, error) {
	return 2, nil
}
func (mockTransactionRepository MockTransactionRepository) GetAverageDebit() (float64, error) {
	return 3, nil
}
