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

func logInfo(msg string) {
	if logLevel <= LoglevelInfo {
		log.Println(fmt.Sprintf("INFO - %s", msg))
	}
}

func logWarn(msg string) {
	if logLevel <= LoglevelWarn {
		log.Println(fmt.Sprintf("WARN - %s", msg))
	}
}

func logDebug(msg string) {
	if logLevel <= LoglevelDebug {
		log.Println(fmt.Sprintf("DEBUG - %s", msg))
	}
}
