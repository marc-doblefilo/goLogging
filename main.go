package main

import (
	"github.com/marc-doblefilo/Logging/logging"
)

func main() {
	logging.SetLevel(logging.INFO)

	logging.Debug("This should not be printed")
	logging.Info("Hello, it works!")
	logging.Warn("This is a warning")
	logging.Error("ERROR!")
}
