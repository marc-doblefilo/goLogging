package logging

import (
	"time"
)

type Level struct {
	name  string
	level int
}

var DEBUG = Level{"DEBUG", 0}
var INFO = Level{"INFO", 10}
var WARN = Level{"WARNING", 20}
var ERROR = Level{"ERROR", 30}
var CRITICAL = Level{"CRITICAL", 40}

var level int = 0

func currentTimeInString() string {
	currentDate := time.Now()

	formattedCurrentDate := currentDate.Format("01-02-2006 15:04.05.0000")

	return formattedCurrentDate
}

func SetLevel(selectedLevel Level) {
	switch selectedLevel {
	case DEBUG:
		level = DEBUG.level
	case INFO:
		level = INFO.level
	case WARN:
		level = WARN.level
	case ERROR:
		level = ERROR.level
	case CRITICAL:
		level = CRITICAL.level
	default:
		return
	}
}
