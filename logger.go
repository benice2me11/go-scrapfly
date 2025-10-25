package scrapfly

import (
	"log"
	"os"
)

// LogLevel defines the severity level for log messages.
type LogLevel int

// Available log levels, from most to least verbose.
const (
	// LevelDebug logs all messages including detailed debug information.
	LevelDebug LogLevel = iota
	// LevelInfo logs informational messages and above.
	LevelInfo
	// LevelWarn logs warnings and errors.
	LevelWarn
	// LevelError logs only error messages.
	LevelError
)

// Logger provides simple leveled logging for the Scrapfly SDK.
type Logger struct {
	logger *log.Logger
	level  LogLevel
}

// NewLogger creates a new Logger instance with the given name prefix.
//
// Example:
//
//	logger := scrapfly.NewLogger("my-scraper")
//	logger.SetLevel(scrapfly.LevelDebug)
//	logger.Info("Starting scraper...")
func NewLogger(name string) *Logger {
	return &Logger{
		logger: log.New(os.Stdout, name+": ", log.LstdFlags),
		level:  LevelInfo,
	}
}

// SetLevel sets the minimum logging level.
// Only messages at this level or higher will be logged.
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a debug-level message.
// These messages are only logged when the level is set to LevelDebug.
func (l *Logger) Debug(v ...interface{}) {
	if l.level <= LevelDebug {
		l.logger.Println(append([]interface{}{"[DEBUG]"}, v...)...)
	}
}

// Info logs an informational message.
// These messages are logged when the level is LevelInfo or lower.
func (l *Logger) Info(v ...interface{}) {
	if l.level <= LevelInfo {
		l.logger.Println(append([]interface{}{"[INFO]"}, v...)...)
	}
}

// Warn logs a warning message.
// These messages are logged when the level is LevelWarn or lower.
func (l *Logger) Warn(v ...interface{}) {
	if l.level <= LevelWarn {
		l.logger.Println(append([]interface{}{"[WARN]"}, v...)...)
	}
}

// Error logs an error message.
// These messages are always logged regardless of the level setting.
func (l *Logger) Error(v ...interface{}) {
	if l.level <= LevelError {
		l.logger.Println(append([]interface{}{"[ERROR]"}, v...)...)
	}
}

// DefaultLogger is the default logger used by the Scrapfly SDK.
//
// You can configure the log level to control verbosity:
//
//	scrapfly.DefaultLogger.SetLevel(scrapfly.LevelDebug)
var DefaultLogger = NewLogger("scrapfly")
