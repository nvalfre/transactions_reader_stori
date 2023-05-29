package transaction_service

import "transactions_reader_stori/domain"

type TransactionServiceI interface {
	ProcessFileContent(fileContent []byte, accountId int, accountName string, email string) error
	GenerateSummary(accountId int) (*domain.SummaryVO, error)
}
