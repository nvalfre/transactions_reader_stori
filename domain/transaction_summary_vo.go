package domain

import "time"

// TransactionSummary represents the transaction specific summary
type TransactionSummary struct {
	ID     int
	Month  time.Time
	Amount float64
}

type MonthlySummary struct {
	Month     time.Month
	Summaries []TransactionSummary
}
