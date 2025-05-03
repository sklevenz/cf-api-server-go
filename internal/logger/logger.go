// Package logger provides a simple structured logger with JSON and text output.
package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Global logger instance
var Log *Logger

type LogLevel int

const (
	// Log levels in increasing order of severity
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

type Logger struct {
	level      LogLevel  // Minimum log level to output
	jsonOut    bool      // Enable JSON formatted output
	fileWriter io.Writer // File output destination (if any)
	stdWriter  io.Writer // Standard output destination
	errWriter  io.Writer // Error output destination
}

// New creates a new Logger instance.
// All logs are written to stdout (e.g., for Docker log collection).
func New(level LogLevel, jsonOut bool, fileWriter2 io.Writer) *Logger {
	log.SetOutput(os.Stdout) // Important: ensure log.Writer() points to stdout

	return &Logger{
		level:      level,
		jsonOut:    jsonOut,
		fileWriter: fileWriter2,
		stdWriter:  os.Stdout,
		errWriter:  os.Stderr,
	}
}

// log writes a log message with the given level, message, and additional fields.
func (l *Logger) log(lvl string, msg string, fields map[string]interface{}) {
	var out string
	if l.jsonOut {
		data := map[string]interface{}{
			"level": lvl,
			"time":  time.Now().Format(time.RFC3339),
			"msg":   msg,
		}
		for k, v := range fields {
			data[k] = v
		}
		b, _ := json.Marshal(data)
		out = string(b)
	} else {
		out = time.Now().Format("2006-01-02 15:04:05") + " [" + lvl + "] " + msg
		if len(fields) > 0 {
			out += " " + fmt.Sprintf("%v", fields)
		}
	}
	if lvl == "ERROR" {
		fmt.Fprintln(l.errWriter, out)

	} else {
		fmt.Fprintln(l.stdWriter, out)
	}
	if l.fileWriter != nil {
		fmt.Fprintln(l.fileWriter, out)
	}
}

// Error logs a message at DEBUG level.
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= LevelDebug {
		l.log("DEBUG", fmt.Sprintf(format, args...), nil)
	}
}

// Error logs a message at INFO level.
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= LevelInfo {
		l.log("INFO", fmt.Sprintf(format, args...), nil)
	}
}

// Error logs a message at WARNING level.
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level <= LevelWarn {
		l.log("WARNING", fmt.Sprintf(format, args...), nil)
	}
}

// Error logs a message at ERROR level.
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= LevelError {
		l.log("ERROR", fmt.Sprintf(format, args...), nil)
	}
}
