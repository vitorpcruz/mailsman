package mailer

import "github.com/mailsman/internal/domain"

type EmailJetMailer interface {
	SendEmail(emailBatch domain.EmailBatch) error
}

type Mailer struct {
}

func (mailer *Mailer) SendEmail(emailBatch domain.EmailBatch) error {
	return nil
}
