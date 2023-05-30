package email_service

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
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
	body := s.buildEmailBodyContent(summary)

	mailer := gomail.NewMessage()
	mailer.SetHeader(from, fmt.Sprintf(senderFormat, s.senderName, s.senderEmail))
	mailer.SetHeader(to, recipient)
	mailer.SetHeader(subject, subjectHealine)
	mailer.SetBody(contentType, body)

	dialer := gomail.NewDialer(s.smtpHost, s.smtpPort, s.smtpUsername, s.smtpPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
