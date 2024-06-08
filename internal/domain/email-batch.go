package domain

import (
	value_object "github.com/mailsman/internal/value_object"
)

type EmailBatch struct {
	Title      value_object.Title
	Message    value_object.Message
	Recipients value_object.Recipients
}

func new(
	title value_object.Title,
	msg value_object.Message,
	recipients value_object.Recipients,
) *EmailBatch {
	return &EmailBatch{Title: title, Message: msg, Recipients: recipients}
}

func NewEmailBatch(
	newTitle, newMsg string,
	emails []string,
) (*EmailBatch, error) {
	title, err := value_object.NewTitle(newTitle)
	if err != nil {
		return nil, err
	}

	msg, err := value_object.NewMessage(newMsg)
	if err != nil {
		return nil, err
	}

	recipients, err := value_object.NewRecipients(emails)
	if len(recipients.EmailAddresses) == 0 && err != nil {
		return nil, err
	}

	return new(title, msg, recipients), err
}
