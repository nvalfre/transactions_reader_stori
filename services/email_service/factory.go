package email_service

// Default SMTP configuration
const (
	DefaultSMTPHost     = "smtp.example.com"
	DefaultSMTPPort     = 587
	DefaultSMTPUsername = "smtpuser"
	DefaultSMTPPassword = "smtppassword"
)

// Default sender details
const (
	DefaultSenderName  = "Your Company"
	DefaultSenderEmail = "noreply@example.com"
)

// NewEmailService creates a new instance of EmailService
func NewEmailServiceDefault() EmailServiceI {
	return NewEmailService(
		DefaultSMTPHost,
		DefaultSMTPPort,
		DefaultSMTPUsername,
		DefaultSMTPPassword,
		DefaultSenderName,
		DefaultSenderEmail,
	)
}

// NewEmailService creates a new instance of EmailService
func NewEmailService(smtpHost string, smtpPort int, smtpUsername, smtpPassword, senderName, senderEmail string) EmailServiceI {
	return &EmailService{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
		senderName:   senderName,
		senderEmail:  senderEmail,
	}
}
