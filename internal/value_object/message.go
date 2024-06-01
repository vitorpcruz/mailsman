package value_object

import (
	"fmt"
	"strings"
)

var (
	MESSAGE_MIN_CHAR = 10
	MESSAGE_MAX_CHAR = 500
)

type Message struct {
	Value string
}

func NewMessage(value string) (Message, error) {
	valueTrimed := strings.TrimSpace(value)
	if err := validateMessage(valueTrimed); err != nil {
		return Message{}, err
	}
	return Message{Value: value}, nil
}

func validateMessage(value string) error {
	if len(value) < MESSAGE_MIN_CHAR || len(value) > MESSAGE_MAX_CHAR || value == "" {
		return fmt.Errorf("the message must be between %v and %v chars", MESSAGE_MIN_CHAR, MESSAGE_MAX_CHAR)
	}
	return nil
}
