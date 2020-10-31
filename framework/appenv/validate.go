package appenv

import (
	"log"
	"os"
	"strings"
)

// Validate validates environment variables before app startup
// Exits/stops the app on validation failure
func Validate() {
	// ------------------------------------
	// Database variables
	// ------------------------------------
	validateNotEmptyF("DATABASE_URL")

}

// GetWithDefault is like os.Getenv, but better
func GetWithDefault(key string, defaultValue string) string {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	}
	return s
}

func validateNotEmptyF(key string) {
	if ok := ValidateNotEmpty(key); !ok {
		log.Fatalf("[env] %s - not found", key)
	}
}

// ValidateNotEmpty validates whether if the value in the environment variable key is empty
func ValidateNotEmpty(key string) bool {
	return strings.TrimSpace(os.Getenv(key)) != ""
}
