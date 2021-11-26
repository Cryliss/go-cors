package log

import (
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
)

// Logger for application logging to the terminal
type Logger struct {
	Log     zerolog.Logger
	Verbose bool
}

// New initializes and returns a new Logger
func New() *Logger {
	// Set the time logging format
	zerolog.TimeFieldFormat = time.RFC3339
	l := Logger{
		Log: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
	return &l
}

// Out prints message to the standard output device and sends it to the logger
func (l *Logger) Out(format string, b ...interface{}) {
	if l.Verbose {
		// Log the message to the terminal in green
		color.HiGreen(format, b...)
	}
}

// OutErr prints message to the standard output error device and sends it to the logger
func (l *Logger) OutErr(format string, b ...interface{}) {
	if l.Verbose {
		// Log the message to the terminal in red
		color.HiRed(format, b...)
	}
}
