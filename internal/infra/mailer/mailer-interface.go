package mailer

import "github.com/mailsman/internal/domain"

type Mailer interface {
	SendEmail(batch domain.EmailBatch) error
}
