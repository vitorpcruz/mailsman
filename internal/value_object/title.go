package value_object

import (
	"fmt"
	"strings"
)

var (
	TITLE_MIN_CHAR = 3
	TITLE_MAX_CHAR = 100
)

type Title struct {
	Value string
}

func NewTitle(value string) (Title, error) {
	valueTrimed := strings.TrimSpace(value)
	if err := validate(valueTrimed); err != nil {
		return Title{Value: ""}, err
	}
	return Title{Value: value}, nil
}

func validate(value string) error {
	if len(value) < TITLE_MIN_CHAR || len(value) > TITLE_MAX_CHAR || value == "" {
		return fmt.Errorf("the title must be between %v and %v chars", TITLE_MIN_CHAR, TITLE_MAX_CHAR)
	}
	return nil
}
