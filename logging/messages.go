package logging

import "fmt"

func Info(message string) {
	if level <= INFO.level {
		fmt.Println(currentTimeInString() + "[" + INFO.name + "]: " + message)
	}
}

func Debug(message string) {
	if level <= DEBUG.level {
		fmt.Println(currentTimeInString() + "[" + DEBUG.name + "]: " + message)
	}
}

func Warn(message string) {
	if level <= WARN.level {
		fmt.Println(currentTimeInString() + "[" + WARN.name + "]: " + message)
	}
}

func Error(message string) {
	if level <= ERROR.level {
		fmt.Println(currentTimeInString() + "[" + ERROR.name + "]: " + message)
	}
}

func Critical(message string) {
	if level <= CRITICAL.level {
		fmt.Println(currentTimeInString() + "[" + CRITICAL.name + "]: " + message)
	}
}
