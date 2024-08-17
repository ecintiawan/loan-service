package notifier

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/pkg/email"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
)

type (
	// repoImpl implements Upload interface
	repoImpl struct {
		emailer email.Email
	}
)

// New creates a new instance of repoImpl
func New(
	emailer email.Email,
) repository.Notifier {
	return &repoImpl{
		emailer: emailer,
	}
}

// Notify will notify related entities based on model
func (r *repoImpl) Notify(
	ctx context.Context,
	model *entity.Notifier,
) error {
	err := r.emailer.Send(email.EmailContent{
		To:      model.To,
		Subject: model.Subject,
		Body:    model.Body,
		Attachment: email.EmailAttachment{
			Content:     model.Attachment.File,
			FileName:    model.Attachment.FileName,
			ContentType: "mime/multipart",
		},
	})
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return nil
}
