package email_service

import "transactions_reader_stori/domain"

type EmailServiceI interface {
	SendSummaryEmail(summary *domain.SummaryVO, recipient string) error
}
