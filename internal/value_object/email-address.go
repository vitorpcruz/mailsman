package value_object

import (
	"net/mail"
)

type EmailAddress struct {
	Value string
}

func NewEmailAddress(value string) (EmailAddress, error) {
	email, err := mail.ParseAddress(value)
	if err != nil {
		return EmailAddress{}, err
	}
	return EmailAddress{Value: email.Address}, err
}
