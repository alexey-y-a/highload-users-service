package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name       string
		inputName  string
		inputEmail string
		wantName   string
		wantEmail  string
		wantErr    error
	}{
		{
			name:       "valid name",
			inputName:  "Alice",
			inputEmail: "alice@example.com",
			wantName:   "Alice",
			wantEmail:  "alice@example.com",
		},
		{
			name:       "trim_spaces_and_lowercase_email",
			inputName:  " Alice ",
			inputEmail: " ALICE@EXAMPLE.COM ",
			wantName:   "Alice",
			wantEmail:  "alice@example.com",
		},
		{
			name:       "cyrillic_name",
			inputName:  "Иванов Иван",
			inputEmail: "ivan@example.com",
			wantName:   "Иванов Иван",
			wantEmail:  "ivan@example.com",
		},
		{
			name:       "empty_name",
			inputName:  "   ",
			inputEmail: "alice@example.com",
			wantErr:    ErrEmptyName,
		},
		{
			name:       "short_name",
			inputName:  "A ",
			inputEmail: "alice@example.com",
			wantErr:    ErrInvalidName,
		},
		{
			name:       "long_name",
			inputName:  strings.Repeat("A", 101),
			inputEmail: "alice@example.com",
			wantErr:    ErrInvalidName,
		},
		{
			name:       "long_cyrillic_name",
			inputName:  strings.Repeat("И", 101),
			inputEmail: "alice@example.com",
			wantErr:    ErrInvalidName,
		},
		{
			name:       "empty_email",
			inputName:  "Alice",
			inputEmail: "    ",
			wantErr:    ErrEmptyEmail,
		},
		{
			name:       "invalid_email",
			inputName:  "Alice",
			inputEmail: "not-a-email",
			wantErr:    ErrInvalidEmail,
		},
		{
			name:       "email_with_display_name",
			inputName:  "Alice",
			inputEmail: "Alice <alice@example.com>",
			wantErr:    ErrInvalidEmail,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user, err := NewUser(test.inputName, test.inputEmail)
			if test.wantErr != nil {
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("NewUser() error = %v, wantErr %v", err, test.wantErr)
				}

				if user != nil {
					t.Fatalf("NewUser() user = %v, expected nil", user)
				}

				return
			}

			if err != nil {
				t.Fatalf("NewUser() unexpected error = %v", err)
			}

			if user.Name != test.wantName {
				t.Fatalf("NewUser() name = %q, want %q", user.Name, test.wantName)
			}

			if user.Email != test.wantEmail {
				t.Fatalf("NewUser() email = %q, want %q", user.Email, test.wantEmail)
			}

			if user.ID != 0 {
				t.Fatalf("NewUser() id = %d, want 0", user.ID)
			}
		})
	}
}
