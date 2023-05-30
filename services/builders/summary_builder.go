package builders

import "transactions_reader_stori/domain"

func BuildSummary(transactionMetadata *domain.TransactionMetadata) *domain.SummaryVO {
	return &domain.SummaryVO{
		TotalBalance:       transactionMetadata.TotalBalance,
		TransactionSummary: transactionMetadata.TransactionSummary,
		MonthlySummary:     transactionMetadata.MonthlySummary,
		AverageCredit:      transactionMetadata.AverageCredit,
		AverageDebit:       transactionMetadata.AverageDebit,
	}
}
