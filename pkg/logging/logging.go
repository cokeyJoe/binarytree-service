// Package logging contains primitive json logger
package logging

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

// Fields log fields
type Fields map[string]interface{}

// Logger json-formatted logger
type Logger struct {
	logger *log.Logger
}

// New returns new logger with dst as destination io.Writer
func New(dst io.Writer) *Logger {

	if dst == nil {
		dst = os.Stdout
	}

	l := log.Logger{}
	l.SetOutput(dst)

	return &Logger{
		logger: &l,
	}
}

// InfoWithFields log fields with level INFO
func (l *Logger) InfoWithFields(fields Fields) {

	fields["level"] = "INFO"
	fields["time"] = time.Now()

	bb, _ := json.Marshal(fields)

	l.logger.Println(string(bb))
}

// ErrorWithFields log fields with level ERROR
func (l *Logger) ErrorWithFields(fields Fields) {

	fields["level"] = "ERROR"
	fields["time"] = time.Now()

	bb, _ := json.Marshal(fields)

	l.logger.Println(string(bb))
}
