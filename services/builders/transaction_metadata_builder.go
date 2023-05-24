package builders

import "transactions_reader_stori/domain"

func BuildTransactionMetadata(totalBalance float64, transactionSummary []domain.TransactionSummary, averageCredit float64, averageDebit float64) *domain.TransactionMetadata {
	return &domain.TransactionMetadata{
		TotalBalance:       totalBalance,
		TransactionSummary: transactionSummary,
		AverageCredit:      averageCredit,
		AverageDebit:       averageDebit,
	}
}
