package domain

// Summary represents the summary information
type Summary struct {
	TotalBalance       float64
	TransactionSummary map[string]int
	AverageCredit      float64
	AverageDebit       float64
}
