package logging

import (
	"time"
)

var DEBUG = "DEBUG"
var INFO = "INFO"
var WARN = "WARN"
var ERROR = "ERROR"
var CRITICAL = "CRITICAL"

var debugLevel = 0
var infoLevel = 10
var warnLevel = 20
var errorLevel = 30
var criticalLevel = 40

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
	case "CRITICAL":
		level = criticalLevel
	default:
		return
	}
}
