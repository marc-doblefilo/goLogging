package logging

import (
	"fmt"
	"time"
)

var DEBUG = "DEBUG"
var INFO = "INFO"
var WARN = "WARN"
var ERROR = "ERROR"

var debugLevel = 0
var infoLevel = 10
var warnLevel = 20
var errorLevel = 30

var level int = 0

func currentTimeInString() string {
	currentDate := time.Now()

	formattedCurrentDate := currentDate.Format("01-02-2006 15:04.05.0000")

	return formattedCurrentDate
}

func SetLevel(Level string) {
	switch Level {
	case "DEBUG":
		level = debugLevel
	case "INFO":
		level = infoLevel
	case "WARN":
		level = warnLevel
	case "ERROR":
		level = errorLevel
	}
}

func Info(message string) {
	if level <= infoLevel {
		fmt.Println(currentTimeInString() + "[" + INFO + "]: " + message)
	}
}

func Debug(message string) {
	if level <= debugLevel {
		fmt.Println(currentTimeInString() + "[" + DEBUG + "]: " + message)
	}
}

func Warn(message string) {
	if level <= warnLevel {
		fmt.Println(currentTimeInString() + "[" + WARN + "]: " + message)
	}
}

func Error(message string) {
	if level <= errorLevel {
		fmt.Println(currentTimeInString() + "[" + ERROR + "]: " + message)
	}
}
