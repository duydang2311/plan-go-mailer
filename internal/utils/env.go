package utils

import (
	"log"
	"os"
)

func RequireEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("require env \"%s\"\n", key)
	}
	return value
}
