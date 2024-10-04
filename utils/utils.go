package utils

import (
	"log"
	"strings"
)

// LogError logs errors with a consistent format
func LogError(err error, message string) {
	if err != nil {
		log.Printf("ERROR: %s - %v", message, err)
	}
}

// CheckStringContains checks if a string contains another string (case-insensitive)
func CheckStringContains(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
