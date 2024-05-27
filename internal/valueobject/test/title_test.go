package valueobject_test

import (
	"strings"
	"testing"

	"github.com/mailsman/internal/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestTitle_MustReturnErrorIfLowerThanMinChar(t *testing.T) {
	title, err := valueobject.NewTitle("")
	assert.Empty(t, title.Value)
	assert.Error(t, err)
}

func TestTitle_MustReturnErrorIfGreaterThanMaxChar(t *testing.T) {
	title, err := valueobject.NewTitle(strings.Repeat("a", 101))
	assert.Empty(t, title.Value)
	assert.Error(t, err)
}

func TestTitle_MustReturnNewTitle(t *testing.T) {
	title, err := valueobject.NewTitle(strings.Repeat("a", 10))
	assert.NotEmpty(t, title.Value)
	assert.Nil(t, err)
}
