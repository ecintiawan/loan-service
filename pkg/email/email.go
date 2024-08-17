package email

import (
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/jordan-wright/email"
)

func NewEmailImpl(config *config.Config) Email {
	return &emailImpl{
		config: config,
	}
}

func (l *emailImpl) Send(content EmailContent) error {
	mail := email.NewEmail()
	mail.To = content.To
	mail.From = fmt.Sprintf("%s <%s>", l.config.Vendor.Email.SenderName, l.config.Credential.Email.SenderEmail)
	mail.Subject = content.Subject
	mail.Text = []byte(content.Body)
	if content.Attachment.Content != nil {
		mail.Attach(
			bytes.NewBuffer(content.Attachment.Content),
			content.Attachment.FileName,
			content.Attachment.ContentType,
		)
	}

	return mail.Send(
		fmt.Sprintf("%s:%s", l.config.Vendor.Email.SMTPHost, l.config.Vendor.Email.SMTPPort),
		smtp.PlainAuth(
			"",
			l.config.Credential.Email.SenderEmail,
			l.config.Credential.Email.SenderPassword,
			l.config.Vendor.Email.SMTPHost,
		),
	)
}
