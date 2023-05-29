package domain

type TransactionMetadata struct {
	TotalBalance       float64
	TransactionSummary []TransactionSummary
	AverageCredit      float64
	AverageDebit       float64
}
