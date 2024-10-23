package main

import (
	"fmt"
	"os"
	"time"
)

// ConsoleLogger struct
type ConsoleLogger struct{}

// Info method for ConsoleLogger
func (cl ConsoleLogger) Info(message string) {
	fmt.Printf("[INFO] %s: %s\n", time.Now().Format(time.RFC3339), message)
}

// Warning method for ConsoleLogger
func (cl ConsoleLogger) Warning(message string) {
	fmt.Printf("[WARNING] %s: %s\n", time.Now().Format(time.RFC3339), message)
}

// Error method for ConsoleLogger
func (cl ConsoleLogger) Error(message string) {
	fmt.Printf("[ERROR] %s: %s\n", time.Now().Format(time.RFC3339), message)
}

// Close method to close the file logger

// FileLogger struct
type FileLogger struct {
	File *os.File
}

// NewFileLogger creates a new FileLogger
func NewFileLogger(filepath string) (*FileLogger, error) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLogger{File: file}, nil
}

// logToFile writes a log message to the file
func (fl *FileLogger) logToFile(level, message string) {
	logMessage := fmt.Sprintf("%s %s: %s\n", time.Now().Format(time.RFC3339), level, message)
	fl.File.WriteString(logMessage)
}

func (fl *FileLogger) Close() {
	fl.File.Close()
}

// Info method for FileLogger
func (fl *FileLogger) Info(message string) {
	fl.logToFile("[INFO]", message)
}

// Warning method for FileLogger
func (fl *FileLogger) Warning(message string) {
	fl.logToFile("[WARNING]", message)
}

// Error method for FileLogger
func (fl *FileLogger) Error(message string) {
	fl.logToFile("[ERROR]", message)
}
