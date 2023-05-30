package email_service

type EmailServiceFactoryI interface {
	NewEmailServiceDefault() EmailServiceI
	newEmailService(smtpHost string, smtpPort int, smtpUsername, smtpPassword, senderName, senderEmail string) EmailServiceI
}

type EmailServiceFactory struct{}

// NewEmailServiceDefault creates a new instance of EmailService
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

// newEmailService creates a new instance of EmailService
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
