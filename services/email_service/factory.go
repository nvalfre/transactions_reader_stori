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

type EmailServiceFactoryI interface {
	NewEmailServiceDefault() EmailServiceI
	newEmailService(smtpHost string, smtpPort int, smtpUsername, smtpPassword, senderName, senderEmail string) EmailServiceI
}

type EmailServiceFactory struct{}

// NewAccountService creates a new instance of AccountService

// NewEmailService creates a new instance of EmailService
func (factory *EmailServiceFactory) NewEmailServiceDefault() EmailServiceI {
	return factory.newEmailService(
		DefaultSMTPHost,
		DefaultSMTPPort,
		DefaultSMTPUsername,
		DefaultSMTPPassword,
		DefaultSenderName,
		DefaultSenderEmail,
	)
}

// NewEmailService creates a new instance of EmailService
func (factory *EmailServiceFactory) newEmailService(smtpHost string, smtpPort int, smtpUsername, smtpPassword, senderName, senderEmail string) EmailServiceI {
	return &EmailService{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
		senderName:   senderName,
		senderEmail:  senderEmail,
	}
}
