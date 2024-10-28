package main

import (
	"fmt"
	"log"
	"os"
)

// Logger interface defines methods for logging
type Logger interface {
	Info(message string)
	Error(message string)
}

// ConsoleLogger struct for logging to the console
type ConsoleLogger struct{}

// FileLogger struct for logging to a file
type FileLogger struct {
	file *os.File
}

// Info method for ConsoleLogger logs information to the console
func (cl ConsoleLogger) Info(message string) {
	log.Println("[INFO]:", message)
}

// Error method for ConsoleLogger logs error messages to the console
func (cl ConsoleLogger) Error(message string) {
	log.Println("[ERROR]:", message)
}

// Info method for FileLogger logs information to a file
func (fl FileLogger) Info(message string) {
	log.SetOutput(fl.file)
	log.Println("[INFO]:", message)
}

// Error method for FileLogger logs error messages to a file
func (fl FileLogger) Error(message string) {
	log.SetOutput(fl.file)
	log.Println("[ERROR]:", message)
}

// A function that uses the logger interface
func ProcessPayment(logger Logger, amount float64) {
	if amount < 0 {
		logger.Error("Invalid payment amount")
	} else {
		logger.Info(fmt.Sprintf("Processed payment of $%.2f", amount))
	}
}

func main() {
	// Create ConsoleLogger
	consoleLogger := ConsoleLogger{}

	// Create and configure FileLogger
	file, err := os.Create("payment.log")
	if err != nil {
		log.Fatal("Unable to create log file:", err)
	}
	defer file.Close()

	fileLogger := FileLogger{file: file}

	// Process payments using ConsoleLogger
	ProcessPayment(consoleLogger, 100.50)
	ProcessPayment(consoleLogger, -10.00)

	// Process payments using FileLogger
	ProcessPayment(fileLogger, 200.00)
	ProcessPayment(fileLogger, -20.00)
}
