package transaction_service

import "transactions_reader_stori/domain"

type TransactionServiceI interface {
	ProcessFile(fileContent []byte) error
	GenerateSummary() (*domain.Summary, error)
}
