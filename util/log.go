package util

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/umardev500/banksampah/constant"
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

func LogParseError(ticket *uuid.UUID, err error, add string) string {
	return fmt.Sprintf("Ticket: %s%v%s | Error: %s. %s%v%s", constant.ANSI_BLUE, ticket, constant.ANSI_RESET, add, constant.ANSI_DARKRED, err, constant.ANSI_RESET)
}
