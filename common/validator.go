package common

import (
	"fmt"
	"net/mail"
	"net/url"
)

func NotEmpty(field, value string) error {
	if len(value) == 0 {
		return fmt.Errorf("%v is empty.", field)
	}

	return nil
}

func ValidEmail(field, value string) error {
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("[%v] %v is not a valid email address.", field, value)
	}

	return nil
}

func ValidUrl(field, value string) error {
	if _, err := url.Parse(value); err != nil {
		return fmt.Errorf("[%v] %v is not a valid url.", field, value)
	}

	return nil
}
