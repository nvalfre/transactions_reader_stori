package email_service

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"sort"
	"strings"
	"time"
	"transactions_reader_stori/domain"
)

const subjectHealine = "SummaryVO Report"

const (
	from         = "From"
	to           = "To"
	subject      = "Subject"
	senderFormat = "%s <%s>"
	contentType  = "text/plain"
)

// SendSummaryEmail sends the summary information as an email
func (s *EmailService) SendSummaryEmail(summary *domain.SummaryVO, recipient string) error {
	body := s.buildBody(summary)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("%s <%s>", s.senderName, s.senderEmail))
	mailer.SetHeader("To", recipient)
	mailer.SetHeader("Subject", subjectHealine)
	mailer.SetBody("text/plain", body)

	dialer := gomail.NewDialer(s.smtpHost, s.smtpPort, s.smtpUsername, s.smtpPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
func (s *EmailService) buildBody(summary *domain.SummaryVO) string {
	sort.Slice(summary.TransactionSummary, func(i, j int) bool {
		return summary.TransactionSummary[i].Month.Before(summary.TransactionSummary[j].Month)
	})

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Total balance: %.2f\n", summary.TotalBalance))

	// Print transaction summaries grouped by month
	prevMonth := time.Time{}
	for _, ts := range summary.TransactionSummary {
		if !ts.Month.Equal(prevMonth) {
			sb.WriteString(fmt.Sprintf("\nMonthly summary for %s:\n", ts.Month.Format("January")))
			prevMonth = ts.Month
		}
		sb.WriteString(fmt.Sprintf("Transaction ID: %d, Amount: %.2f\n", ts.ID, ts.Amount))
	}

	sb.WriteString(fmt.Sprintf("\nAverage credit amount: %.2f\n", summary.AverageCredit))
	sb.WriteString(fmt.Sprintf("Average debit amount: %.2f\n", summary.AverageDebit))

	body := sb.String()
	return body
}
