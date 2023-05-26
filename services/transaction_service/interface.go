package transaction_service

import "transactions_reader_stori/domain"

type TransactionServiceI interface {
	ProcessFileContent(fileContent []byte, accountId string, accountName string) error
	GenerateSummary(accountId string) (*domain.SummaryVO, error)
}
