package utils

import (
	"errors"
	"regexp"
	"strings"
)

func IsValidEmail(email string) error {
	// This is a simple email validation regex; you may use a more comprehensive one if needed
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return errors.New("Invalid email formatting")
	}

	if !strings.HasSuffix(email, "@iiitl.ac.in") {
		return errors.New("Not a IIITL email")
	}
	return nil
}

func IsValidPassword(password string) error {
	// Password must be at least 8 characters long and contain at least:
	// - One uppercase letter
	// - One lowercase letter
	// - One digit
	// - One special character (e.g., !@#$%^&*)
	if len(password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("Password must contain at least one uppercase letter")
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("Password must contain at least one lowercase letter")
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return errors.New("Password must contain at least one digit")
	}

	if !regexp.MustCompile(`[@#$%^&*!]`).MatchString(password) {
		return errors.New("Password must contain at least one special character (e.g., !@#$%^&*)")
	}

	return nil
}
