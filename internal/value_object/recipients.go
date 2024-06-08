package value_object

import "errors"

type Recipients struct {
	EmailAddresses []EmailAddress
}

func NewRecipients(emails []string) (Recipients, error) {
	emailAddresses := []EmailAddress{}
	errList := []error{}
	for _, e := range emails {
		email, err := NewEmailAddress(e)

		if err != nil {
			errList = append(errList, err)
		} else {
			emailAddresses = append(emailAddresses, email)
		}
	}

	err := errors.Join(errList...)

	return Recipients{
		EmailAddresses: emailAddresses,
	}, err
}
