package log

import (
	"fmt"
)

const (
	DebugLevel = iota
	InfoLevel  = iota
	WarnLevel  = iota
	ErrorLevel = iota
	FatalLevel = iota
)

type Logger struct {
	MinLevel int
	Levels   [5]string
}

func NewLogger(minLevel int) *Logger {
	return &Logger{
		Levels:   [5]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"},
		MinLevel: minLevel,
	}
}

func (l *Logger) Debug(message string) {
	l.log(DebugLevel, message)
}

func (l *Logger) Info(message string) {
	l.log(InfoLevel, message)
}

func (l *Logger) Warn(message string) {
	l.log(WarnLevel, message)
}

func (l *Logger) Error(message string) {
	l.log(ErrorLevel, message)
}

func (l *Logger) Fatal(message string) {
	l.log(FatalLevel, message)
}

func (l *Logger) log(level int, message string) {
	if level < l.MinLevel {
		return
	}
	fmt.Printf("[%v] %v\n", l.Levels[level], message)
}
