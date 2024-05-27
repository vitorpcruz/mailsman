package valueobject

import (
	"strings"
	"testing"

	"github.com/mailsman/internal/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestMessage_MustReturnErrorIfLowerThanMinChars(t *testing.T) {
	v := strings.Repeat("a", valueobject.MESSAGE_MIN_CHAR-1)

	message, err := valueobject.NewMessage(v)
	assert.Nil(t, message)
	assert.Error(t, err)
}

func TestMessage_MustReturnErrorIfGreaterThanMaxChars(t *testing.T) {
	v := strings.Repeat("a", valueobject.MESSAGE_MAX_CHAR+1)

	message, err := valueobject.NewMessage(v)
	assert.Nil(t, message)
	assert.Error(t, err)
}

func TestMessage_MustReturnValidMessage(t *testing.T) {
	v := strings.Repeat("a", valueobject.MESSAGE_MAX_CHAR)

	message, err := valueobject.NewMessage(v)
	assert.NotNil(t, message)
	assert.Nil(t, err)
	assert.Equal(t, v, message.Value)
}
