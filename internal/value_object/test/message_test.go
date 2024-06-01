package value_object

import (
	"strings"
	"testing"

	"github.com/mailsman/internal/value_object"
	"github.com/stretchr/testify/assert"
)

func TestMessage_MustReturnErrorIfLowerThanMinChars(t *testing.T) {
	v := strings.Repeat("a", value_object.MESSAGE_MIN_CHAR-1)

	message, err := value_object.NewMessage(v)
	assert.Empty(t, message)
	assert.Error(t, err)
}

func TestMessage_MustReturnErrorIfGreaterThanMaxChars(t *testing.T) {
	v := strings.Repeat("a", value_object.MESSAGE_MAX_CHAR+1)

	message, err := value_object.NewMessage(v)
	assert.Empty(t, message)
	assert.Error(t, err)
}

func TestMessage_MustReturnValidMessage(t *testing.T) {
	v := strings.Repeat("a", value_object.MESSAGE_MAX_CHAR)

	message, err := value_object.NewMessage(v)
	assert.NotEmpty(t, message)
	assert.Empty(t, err)
	assert.Equal(t, v, message.Value)
}
