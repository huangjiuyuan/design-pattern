package logger

import (
	"fmt"
)

const (
	_ = iota
	// INFO level logger.
	INFO
	// DEBUG level logger.
	DEBUG
	// ERROR level logger.
	ERROR
)

type AbstractLogger interface {
	SetNextLogger(logger AbstractLogger)
	LogMessage(level int, message string)
}

type logger struct {
	Level      int
	NextLogger AbstractLogger
	Write      func(message string)
}

func (l *logger) SetNextLogger(logger AbstractLogger) {
	l.NextLogger = logger
}

func (l *logger) LogMessage(level int, message string) {
	if l.Level <= level {
		l.Write(message)
	}
	if l.NextLogger != nil {
		l.NextLogger.LogMessage(level, message)
	}
}

type ConsoleLogger struct {
	logger
}

func NewConsoleLogger() *ConsoleLogger {
	cl := new(ConsoleLogger)
	cl.Level = INFO
	cl.Write = func(message string) {
		fmt.Printf("Info message: %s\n", message)
	}
	return cl
}

type FileLogger struct {
	logger
}

func NewFileLogger() *FileLogger {
	fl := new(FileLogger)
	fl.Level = DEBUG
	fl.Write = func(message string) {
		fmt.Printf("Debug message: %s\n", message)
	}
	return fl
}

type ErrorLogger struct {
	logger
}

func NewErrorLogger() *ErrorLogger {
	el := new(ErrorLogger)
	el.Level = ERROR
	el.logger.Write = func(message string) {
		fmt.Printf("Error message: %s\n", message)
	}
	return el
}

func LoggerChain() AbstractLogger {
	cl := NewConsoleLogger()
	fl := NewFileLogger()
	el := NewErrorLogger()

	el.SetNextLogger(fl)
	fl.SetNextLogger(cl)
	return el
}
