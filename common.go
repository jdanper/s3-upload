package main

import (
	"fmt"
	"log"
	"os"
)

var env = envOrDefault("ENV", "debug")

func envOrDefault(envVar, defVal string) string {
	envVal := os.Getenv(envVar)
	if envVal == "" {
		return defVal
	}

	return envVal
}

func fatalEnv(key string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		log.Fatalln(fmt.Sprintf("environment variable %s not set.", key))
	}

	return envValue
}