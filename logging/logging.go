package logging

import (
	"fmt"
	"time"
)

var DEBUG = "DEBUG"
var INFO = "INFO"
var WARN = "WARN"
var ERROR = "ERROR"

func currentTimeInString() string {
	currentDate := time.Now()

	formattedCurrentDate := currentDate.Format("01-02-2006 15:04.05.0000")

	return formattedCurrentDate
}

func Info(message string) {
	fmt.Println(currentTimeInString() + "[" + INFO + "]: " + message)
}

func Debug(message string) {
	fmt.Println(currentTimeInString() + "[" + DEBUG + "]: " + message)
}
