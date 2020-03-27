package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	LoglevelDebug = iota + 1
	LoglevelInfo
	LoglevelWarn
	LoglevelError
)

var logLevel, _ = strconv.Atoi(envOrDefault("LOG_LEVEL", "1"))

func setLogLevel(level int) {
	if level != 0 {
		logLevel = level
	}
}

func logError(msg string) {
	if logLevel <= LoglevelError {
		log.Println(fmt.Sprintf("ERROR - %s", msg))
	}
}