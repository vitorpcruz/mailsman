package valueobject

import (
	"net/mail"
)

type EmailAddress struct {
	Address string
}

func NewEmailAddress(value string) (*EmailAddress, error) {
	email, err := mail.ParseAddress(value)
	if err != nil {
		return nil, err
	}
	return &EmailAddress{Address: email.Address}, err
}
