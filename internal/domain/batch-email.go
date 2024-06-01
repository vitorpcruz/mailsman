package domain

import (
	"errors"

	value_object "github.com/mailsman/internal/value_object"
)

type BatchEmail struct {
	Title          value_object.Title
	Message        value_object.Message
	EmailAddresses []value_object.EmailAddress
}

func newBatchEmail(
	title value_object.Title,
	msg value_object.Message,
	emailAddresses []value_object.EmailAddress,
) *BatchEmail {
	return &BatchEmail{Title: title, Message: msg, EmailAddresses: emailAddresses}
}

func NewBatchEmail(
	newTitle, newMsg string,
	emails []string,
) (*BatchEmail, error) {
	title, err := value_object.NewTitle(newTitle)
	if err != nil {
		return nil, err
	}

	msg, err := value_object.NewMessage(newMsg)
	if err != nil {
		return nil, err
	}

	emailAddresses := []value_object.EmailAddress{}
	errList := []error{}
	for _, e := range emails {
		email, err := value_object.NewEmailAddress(e)

		if err != nil {
			errList = append(errList, err)
		} else {
			emailAddresses = append(emailAddresses, email)
		}
	}

	err = errors.Join(errList...)

	if len(emailAddresses) == 0 && len(errList) > 0 {
		return nil, err
	}

	return newBatchEmail(title, msg, emailAddresses), err
}
