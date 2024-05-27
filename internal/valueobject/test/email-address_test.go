package valueobject_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/mailsman/internal/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestEmailAddress_MustReturnErrorIfInvalidEmail(t *testing.T) {
	emails := []string{"abc", "bc@", "@pirilampos.com"}

	for _, e := range emails {
		email, err := valueobject.NewEmailAddress(e)
		assert.Nil(t, email)
		assert.Error(t, err)
	}
}

func TestEmailAddress_MustReturnValidEmail(t *testing.T) {
	emails := []string{}

	for i := 0; i < 10; i++ {
		emails = append(emails, fake.EmailAddress())
	}

	for _, e := range emails {
		email, err := valueobject.NewEmailAddress(e)
		assert.NotNil(t, email)
		assert.Nil(t, err)
		assert.Equal(t, email.Address, e)
	}
}
