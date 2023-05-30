package domain

import "time"

type TransactionMetadata struct {
	TotalBalance       float64
	TransactionSummary []TransactionSummary
	MonthlySummary     map[time.Month]MonthlySummary
	AverageCredit      float64
	AverageDebit       float64
}
