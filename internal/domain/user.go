package domain

import (
	"errors"
	"net/mail"
	"strings"
	"unicode/utf8"
)

const (
	minNameLength  = 2
	maxNameLength  = 100
	maxEmailLength = 255
)

var (
	ErrEmptyName    = errors.New("name is required")
	ErrInvalidName  = errors.New("name must be between 2 and 100 characters")
	ErrEmptyEmail   = errors.New("email is required")
	ErrInvalidEmail = errors.New("email is invalid")
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))

	if name == "" {
		return nil, ErrEmptyName
	}

	nameLength := utf8.RuneCountInString(name)

	if nameLength < minNameLength {
		return nil, ErrInvalidName
	}

	if nameLength > maxNameLength {
		return nil, ErrInvalidName
	}

	if email == "" {
		return nil, ErrEmptyEmail
	}

	if len(email) > maxEmailLength {
		return nil, ErrInvalidEmail
	}

	address, err := mail.ParseAddress(email)
	if err != nil {
		return nil, ErrInvalidEmail
	}

	if address.Address != email {
		return nil, ErrInvalidEmail
	}

	return &User{
		Name:  name,
		Email: email,
	}, nil
}
