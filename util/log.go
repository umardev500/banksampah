package util

import (
	"fmt"
	"sync"
)

type Logger struct{}

var (
	loggerInstance *Logger
	once           sync.Once
)

func NewLogger() *Logger {
	once.Do(func() {
		loggerInstance = &Logger{}
	})

	return loggerInstance
}

func (l *Logger) FirstLine() {
	fmt.Print("\r")
}

// Move the cursor up one line
func (l *Logger) Upline() {
	fmt.Print("\033[1A")
}

// Clear the previous message
func (l *Logger) ClearPrev() {
	fmt.Print("\033[2K")
}

// Move up and clear the previous message
func (l *Logger) UplineClearPrev() {
	l.Upline()
	l.ClearPrev()
}
