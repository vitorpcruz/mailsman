package value_object_test

import (
	"strings"
	"testing"

	"github.com/mailsman/internal/value_object"
	"github.com/stretchr/testify/assert"
)

func TestTitle_MustReturnErrorIfLowerThanMinChar(t *testing.T) {
	title, err := value_object.NewTitle("")
	assert.Empty(t, title.Value)
	assert.Error(t, err)
}

func TestTitle_MustReturnErrorIfGreaterThanMaxChar(t *testing.T) {
	title, err := value_object.NewTitle(strings.Repeat("a", 101))
	assert.Empty(t, title.Value)
	assert.Error(t, err)
}

func TestTitle_MustReturnNewTitle(t *testing.T) {
	title, err := value_object.NewTitle(strings.Repeat("a", 10))
	assert.NotEmpty(t, title.Value)
	assert.Nil(t, err)
}
