package value_object_test

import (
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mailsman/internal/value_object"
	"github.com/stretchr/testify/assert"
)

func TestEmailAddress_MustReturnErrorIfInvalidEmail(t *testing.T) {
	email, err := value_object.NewEmailAddress(
		strings.Replace(
			faker.Email(), "@", "|", 1,
		),
	)

	assert.Empty(t, email)
	assert.Error(t, err)
}

func TestEmailAddress_MustReturnValidEmail(t *testing.T) {
	e := faker.Email()
	email, err := value_object.NewEmailAddress(e)

	assert.NotEmpty(t, email)
	assert.Nil(t, err)
	assert.Equal(t, e, email.Value)
}
