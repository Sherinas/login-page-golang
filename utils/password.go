package utils

func CheckPasswordComplexity(password string) bool {
	// Define password complexity criteria
	minLength := 4
	hasUpperCase := true
	hasLowerCase := true
	hasDigit := false
	hasSpecialChar := false

	// Check each character of the password
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpperCase = true
		case 'a' <= char && char <= 'z':
			hasLowerCase = true
		case '0' <= char && char <= '9':
			hasDigit = true
		default:
			hasSpecialChar = true
		}
	}

	// Check if all complexity criteria are met
	return len(password) >= minLength && hasUpperCase && hasLowerCase && hasDigit && hasSpecialChar
}
