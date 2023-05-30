package builders

import (
	"time"
	"transactions_reader_stori/domain"
)

func BuildTransactionMetadata(totalBalance float64, transactionSummary []domain.TransactionSummary, averageCredit float64, averageDebit float64) *domain.TransactionMetadata {
	return &domain.TransactionMetadata{
		TotalBalance:       totalBalance,
		TransactionSummary: transactionSummary,
		MonthlySummary:     groupTransactionSummariesByMonth(transactionSummary),
		AverageCredit:      averageCredit,
		AverageDebit:       averageDebit,
	}
}

func groupTransactionSummariesByMonth(transactionSummaries []domain.TransactionSummary) map[time.Month]domain.MonthlySummary {
	monthlySummaries := make(map[time.Month]domain.MonthlySummary)

	for _, summary := range transactionSummaries {
		month := summary.Month.Month()

		monthlySummary, exists := monthlySummaries[month]
		if !exists {
			monthlySummary = domain.MonthlySummary{
				Month:     month,
				Summaries: []domain.TransactionSummary{summary},
			}
		} else {
			monthlySummary.Summaries = append(monthlySummary.Summaries, summary)
		}

		monthlySummaries[month] = monthlySummary
	}

	return monthlySummaries
}
