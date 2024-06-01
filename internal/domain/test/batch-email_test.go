package domain_test

import (
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mailsman/internal/domain"
	"github.com/mailsman/internal/value_object"
	"github.com/stretchr/testify/assert"
)

func TestBatchEmail_NewBatchEmail_InvalidTitle_MustReturnError(t *testing.T) {
	batch, err := domain.NewBatchEmail(
		"aa",
		"this is a messsage",
		[]string{
			"nutts@nuttsc.com",
		},
	)
	assert.Nil(t, batch)
	assert.Error(t, err)
}

func TestBatchEmail_NewBatchEmail_InvalidMessage_MustReturnError(t *testing.T) {
	batch, err := domain.NewBatchEmail(
		"This is a title",
		strings.Repeat("a", value_object.TITLE_MIN_CHAR-1),
		[]string{
			"nutts@nuttsc.com",
		},
	)
	assert.Nil(t, batch)
	assert.Error(t, err)
}

func TestBatchEmail_NewBatchEmail_AllEmailsInvalid_MustReturnError(t *testing.T) {
	emails := []string{}
	for i := 0; i < 5; i++ {
		email := strings.Replace(faker.Email(), "@", "||", 1)
		emails = append(emails, email)
	}

	batch, err := domain.NewBatchEmail(
		strings.Repeat("a", value_object.MESSAGE_MIN_CHAR),
		strings.Repeat("a", value_object.TITLE_MIN_CHAR),
		emails,
	)

	assert.Nil(t, batch)
	assert.Error(t, err)
}

func TestBatchEmail_NewBatchEmail_UniqueEmailValid_MustReturnBatchAndError(t *testing.T) {
	emailList := make([]string, 0)

	for i := 0; i < 5; i++ {
		email := strings.Replace(faker.Email(), "@", "||", 1)
		emailList = append(emailList, email)
	}

	email := faker.Email()

	emailList = append(emailList, email)

	batch, err := domain.NewBatchEmail(
		strings.Repeat("a", value_object.TITLE_MAX_CHAR),
		strings.Repeat("a", value_object.MESSAGE_MAX_CHAR),
		emailList,
	)

	assert.NotNil(t, batch)

	assert.Error(t, err)

	assert.Equal(t, len(batch.EmailAddresses), 1)

	assert.Equal(t, batch.EmailAddresses[0].Value, email)
}
