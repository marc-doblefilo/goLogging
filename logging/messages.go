package logging

import "fmt"

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

func Critical(message string) {
	if level <= criticalLevel {
		fmt.Println(currentTimeInString() + "[" + CRITICAL + "]: " + message)
	}
}
