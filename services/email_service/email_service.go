package email_service

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"transactions_reader_stori/domain"
)

const subject = "SummaryVO Report"
const emailBody = `Total balance: %.2f
				Number of transactions in July: %d
				Number of transactions in August: %d
				Average credit amount: %.2f
				Average debit amount: %.2f`

// SendSummaryEmail sends the summary information as an email
func (s *EmailService) SendSummaryEmail(summary *domain.SummaryVO, recipient string) error {
	body := fmt.Sprintf(emailBody, summary.TotalBalance, summary.TransactionSummary[0].Amount, summary.TransactionSummary[1].Amount, summary.AverageCredit, summary.AverageDebit)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("%s <%s>", s.senderName, s.senderEmail))
	mailer.SetHeader("To", recipient)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	dialer := gomail.NewDialer(s.smtpHost, s.smtpPort, s.smtpUsername, s.smtpPassword)

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
