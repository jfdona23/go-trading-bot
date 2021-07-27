package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	// ANSI color definitions for terminal outputs.
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"

	// Log level prefixes. Colored with ANSI color definitions.
	infoColored  = colorWhite + "INFO " + colorReset
	warnColored  = colorYellow + "WARN " + colorReset
	errorColored = colorRed + "ERROR" + colorReset
	debugColored = colorCyan + "DEBUG" + colorReset

	// Log levels.
	_DEBUG  = 0
	_ERROR  = 1
	_WARN   = 2
	_INFO   = 3
	_DEFLVL = _INFO

	// Default time format for logging.
	timeFormat = "2006-01-02 15:04:05"
)

// The global logger variable to export across the package.
var Log LoggerInterface

// Interface for new logger objects.
type LoggerInterface interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	GetLevel() string
	SetLevel(level string)
	AllLevels() map[int]string
}

// The logger struct that conforms the Logger interface.
type Logger struct {
	LogLevel int // Minimum log level to print (0:debug 1:error 2:warn 3:info).
}

// Get the current log level.
func (l *Logger) GetLevel() string {
	return l.AllLevels()[l.LogLevel]
}

// Set the log level.
func (l *Logger) SetLevel(level string) {
	level = strings.ToLower(level)
	for levelInt, levelStr := range l.AllLevels() {
		if level == levelStr {
			l.LogLevel = levelInt
			return
		}
	}
	l.LogLevel = _DEFLVL
}

// Returns a map with all the current allowed log levels.
//
// The keys represent the proper log level while the values
// are the string representation of it.
func (l *Logger) AllLevels() map[int]string {
	return map[int]string{0: "debug", 1: "error", 2: "warn", 3: "info"}
}

// Implementation for the INFO log level.
func (l *Logger) Info(args ...interface{}) {
	if _INFO >= l.LogLevel {
		fmt.Printf(infoColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

// Implementation for the WARN log level.
func (l *Logger) Warn(args ...interface{}) {
	if _WARN >= l.LogLevel {
		fmt.Printf(warnColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

// Implementation for the ERROR log level.
func (l *Logger) Error(args ...interface{}) {
	if _ERROR >= l.LogLevel {
		fmt.Printf(errorColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

// Implementation for the DEBUG log level.
func (l *Logger) Debug(args ...interface{}) {
	if _DEBUG >= l.LogLevel {
		fmt.Printf(debugColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

// SetLogger is the setter for log variable, it should be the only way to assign value to log.
func SetLogger(newLogger LoggerInterface) {
	Log = newLogger
}
