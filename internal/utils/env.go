package utils

import "os"

// GetEnv looks up the value of an environment variable based on a specified key.
// If the environment variable with the given key exists, the function returns its value.
// Otherwise, it returns a fallback value provided as an argument.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
