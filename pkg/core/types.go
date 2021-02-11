package core

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Email struct {
	From string
	To   string

	Subject string
	Content string
}

func (e Email) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.From, validation.Required, is.Email),
		validation.Field(&e.To, validation.Required, is.Email),
	)
}

type Filter interface {
	// Test checks a mail and returns true if the mail should trigger a notification
	Test(email Email) bool
}
