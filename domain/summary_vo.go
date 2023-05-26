package domain

// SummaryVO represents the summary information
type SummaryVO struct {
	TotalBalance       float64
	TransactionSummary []TransactionSummary
	AverageCredit      float64
	AverageDebit       float64
}
