package domain

import "time"

// SummaryVO represents the summary information
type SummaryVO struct {
	TotalBalance       float64
	TransactionSummary []TransactionSummary
	MonthlySummary     map[time.Month]MonthlySummary
	AverageCredit      float64
	AverageDebit       float64
}
