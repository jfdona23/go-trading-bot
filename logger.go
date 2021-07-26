package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"

	infoColored  = colorWhite + "INFO " + colorReset
	warnColored  = colorYellow + "WARN " + colorReset
	errorColored = colorRed + "ERROR" + colorReset
	debugColored = colorCyan + "DEBUG" + colorReset

	DEBUG = 0
	ERROR = 1
	WARN  = 2
	INFO  = 3
	// Default log level.
	DEFLVL = INFO

	// Default time format for logging.
	timeFormat = "2006-01-02 15:04:05"
)

// The global logger variable to export across the package.
var Log Logger

// Interface for new logger objects.
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	getLevel() string
	setLevel(level string)
	allLevels() map[int]string
}

// The logger struct that conforms the Logger interface.
type logger struct {
	logLevel int // Minimum log level to print (0:debug 1:error 2:warn 3:info).
}

// Get the current log level.
func (l *logger) getLevel() string {
	return l.allLevels()[l.logLevel]
}

// Set the log level.
func (l *logger) setLevel(level string) {
	level = strings.ToLower(level)
	for levelInt, levelStr := range l.allLevels() {
		if level == levelStr {
			l.logLevel = levelInt
			return
		}
	}
	l.logLevel = DEFLVL
}

// Returns a map with all the current allowed log levels.
//
// The keys represent the proper log level while the values
// are the string representation of it.
func (l *logger) allLevels() map[int]string {
	return map[int]string{0: "debug", 1: "error", 2: "warn", 3: "info"}
}

func (l *logger) Info(args ...interface{}) {
	if INFO >= l.logLevel {
		fmt.Printf(infoColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

func (l *logger) Warn(args ...interface{}) {
	if WARN >= l.logLevel {
		fmt.Printf(warnColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

func (l *logger) Error(args ...interface{}) {
	if ERROR >= l.logLevel {
		fmt.Printf(errorColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

func (l *logger) Debug(args ...interface{}) {
	if DEBUG >= l.logLevel {
		fmt.Printf(debugColored+" -- "+time.Now().Format(timeFormat)+" -- %s \n", args...)
	}
}

// setLogger is the setter for log variable, it should be the only way to assign value to log.
func setLogger(newLogger Logger) {
	Log = newLogger
}
