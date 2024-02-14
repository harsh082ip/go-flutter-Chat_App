package helpers

import (
	"regexp"
)

func CheckPasswordValidity(password string) bool {
	// 1st check - Length
	if len(password) < 8 {
		return false
	}

	// 2nd check - Lowercase, Uppercase, Special characters, and at least one numeric digit
	var hasLower, hasUpper, hasSpecial, hasDigit bool
	for i := 0; i < len(password); i++ {
		char := string(password[i])
		switch {
		case 'a' <= password[i] && password[i] <= 'z':
			hasLower = true
		case 'A' <= password[i] && password[i] <= 'Z':
			hasUpper = true
		case regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(char):
			hasSpecial = true
		case '0' <= password[i] && password[i] <= '9':
			hasDigit = true
		}
	}

	// 3rd check - Special Characters
	if !hasSpecial {
		return false
	}

	// 4th check - Lowercase, Uppercase, Special character, and at least one numeric digit should be present
	return hasLower && hasUpper && hasDigit
}
