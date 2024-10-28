package main

import (
	"fmt"
)

func main() {
	// Using ConsoleLogger
	consoleLogger := ConsoleLogger{}
	consoleLogger.Info("This is an info message.")
	consoleLogger.Warning("This is a warning message.")
	consoleLogger.Error("This is an error message.")

	// Using FileLogger
	fileLogger, err := NewFileLogger("app.log")
	if err != nil {
		fmt.Println("Error creating file logger:", err)
		return
	}
	defer fileLogger.Close()

	fileLogger.Info("This is an info message in the file.")
	fileLogger.Warning("This is a warning message in the file.")
	fileLogger.Error("This is an error message in the file.")
}
