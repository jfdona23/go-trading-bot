package helpers

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Type Alias
type LoggerType = *logrus.Logger

var (
	logLevel  string                  = Getenv("LOG", "info")
	logLevels map[string]logrus.Level = map[string]logrus.Level{
		"info":  logrus.InfoLevel,
		"error": logrus.ErrorLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
	}
)

// Common logging configuration
func GetLogger() LoggerType {
	var log = logrus.New()
	logLevel = strings.ToLower(logLevel)
	logFormat := new(logrus.TextFormatter)
	logFormat.TimestampFormat = "2006-01-02 15:04:05"
	logFormat.FullTimestamp = true
	log.SetFormatter(logFormat)
	log.SetLevel(logLevels[logLevel])
	return log
}

// Get an environment value or returns a fallback value if not present
func Getenv(environmentVar string, defaultValue string) string {
	value, exists := os.LookupEnv(environmentVar)
	if !exists {
		return defaultValue
	}
	return value
}
