package email_service

// EmailService handles sending summary emails
type EmailService struct {
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
	senderName   string
	senderEmail  string
}
