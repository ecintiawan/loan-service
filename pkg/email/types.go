package email

import "github.com/ecintiawan/loan-service/pkg/config"

type (
	Email interface {
		Send(content EmailContent) error
	}

	EmailContent struct {
		To         []string
		Subject    string
		Body       string
		Attachment EmailAttachment
	}

	EmailAttachment struct {
		Content     []byte
		FileName    string
		ContentType string
	}

	emailImpl struct {
		config *config.Config
	}
)
