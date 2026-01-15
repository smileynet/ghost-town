package logger

import (
	"fmt"
	"os"
)

// Logger interface for logging
type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
}

// DefaultLogger is a simple console logger
type DefaultLogger struct{}

func (l *DefaultLogger) Info(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

func (l *DefaultLogger) Debug(msg string) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("[DEBUG] %s\n", msg)
	}
}

func (l *DefaultLogger) Error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}
