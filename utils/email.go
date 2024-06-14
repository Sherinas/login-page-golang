package utils

import (
	"regexp"

	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
	DB = db
}

func ValidateEmailAddress(email string) bool {
	// Regular expression for basic email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
