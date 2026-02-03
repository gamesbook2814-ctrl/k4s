package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Level represents log level
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger writes logs to ~/.k4s/logs directory
type Logger struct {
	mu       sync.Mutex
	file     *os.File
	minLevel Level
	enabled  bool
}

var (
	defaultLogger *Logger
	once          sync.Once
)

// Init initializes the default logger
func Init(minLevel Level) error {
	var initErr error
	once.Do(func() {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			initErr = fmt.Errorf("get home dir: %w", err)
			return
		}

		logDir := filepath.Join(homeDir, ".k4s", "logs")
		if err := os.MkdirAll(logDir, 0755); err != nil {
			initErr = fmt.Errorf("create log dir: %w", err)
			return
		}

		// Create log file with date
		logFile := filepath.Join(logDir, fmt.Sprintf("k4s-%s.log", time.Now().Format("2006-01-02")))
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			initErr = fmt.Errorf("open log file: %w", err)
			return
		}

		defaultLogger = &Logger{
			file:     file,
			minLevel: minLevel,
			enabled:  true,
		}

		// Write startup message
		defaultLogger.log(LevelInfo, "k4s logger initialized")
	})
	return initErr
}

// Close closes the logger
func Close() {
	if defaultLogger != nil && defaultLogger.file != nil {
		defaultLogger.file.Close()
	}
}

// SetEnabled enables or disables logging
func SetEnabled(enabled bool) {
	if defaultLogger != nil {
		defaultLogger.enabled = enabled
	}
}

// log writes a log message
func (l *Logger) log(level Level, format string, args ...interface{}) {
	if l == nil || !l.enabled || level < l.minLevel {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	msg := fmt.Sprintf(format, args...)
	line := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level.String(), msg)

	l.file.WriteString(line)
}

// Debug logs a debug message
func Debug(format string, args ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.log(LevelDebug, format, args...)
	}
}

// Info logs an info message
func Info(format string, args ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.log(LevelInfo, format, args...)
	}
}

// Warn logs a warning message
func Warn(format string, args ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.log(LevelWarn, format, args...)
	}
}

// Error logs an error message
func Error(format string, args ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.log(LevelError, format, args...)
	}
}

// Errorf logs an error with the error object
func Errorf(err error, format string, args ...interface{}) {
	if defaultLogger != nil {
		msg := fmt.Sprintf(format, args...)
		defaultLogger.log(LevelError, "%s: %v", msg, err)
	}
}
