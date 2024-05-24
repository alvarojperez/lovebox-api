package main

import (
	"log"
	"os"
)

func getEnvVariable(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}

	return value
}