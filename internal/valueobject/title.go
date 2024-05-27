package valueobject

import (
	"fmt"
	"strings"
)

var (
	title_min_char = 3
	title_max_char = 100
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
	if len(value) < title_min_char || len(value) > title_max_char || value == "" {
		return fmt.Errorf("the title must be between %v and %v chars", title_min_char, title_max_char)
	}
	return nil
}
