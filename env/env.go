package env

import (
	"os"
)

func Get(key string, def string) string {
	result := os.Getenv(key)
	if result == "" {
		result = def
	}

	return result
}
