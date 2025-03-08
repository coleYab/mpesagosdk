package logger

import (
	"os"
	"strings"
	"log/slog"
)

type LogLevel int

const (
	DEBUG LogLevel = iota // Logs detailed debugging information.
	INFO                  // Logs general information messages.
	WARN                  // Logs warning messages indicating potential issues.
	ERROR                 // Logs error messages indicating failures or significant issues.
)

type Logger struct {
	level LogLevel
    logger *slog.Logger
}

func NewLogger(level LogLevel) *Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(handler)

	return &Logger{
		level:  level,
		logger: logger,
	}
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.level <= DEBUG {
		l.logger.Debug(msg, args...)
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.level <= INFO {
		l.logger.Info(msg, args...)
	}
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	if l.level <= WARN {
		l.logger.Warn(msg, args...)
	}
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.level <= ERROR {
		l.logger.Error(msg, args...)
	}
}

func ParseLevel(level string) LogLevel {
	switch strings.ToUpper(level) {
	case "INFO":
		return INFO
    case "DEBUG":
        return DEBUG
	case "ERROR":
		return ERROR
	default:
        return WARN
	}
}

