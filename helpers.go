package main

import "os"

// Get an environment value or returns a fallback value if not present
func getenv(environmentVar string, defaultValue string) string {
	value, exists := os.LookupEnv(environmentVar)
	if !exists {
		return defaultValue
	}
	return value
}
