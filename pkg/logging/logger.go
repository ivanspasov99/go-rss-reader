package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

var logger *logWrapper

// zerolog is a struct that wraps the zerolog logger
type logWrapper struct {
	log *zerolog.Logger
}

// GetLogger initialize zerolog or retrieves it if it was already initialized
func GetLogger() *logWrapper {
	if logger != nil {
		fmt.Println("initialized already")
		return logger
	}
	fmt.Println("not initialized")
	l := zerolog.New(os.Stdout)
	logger = &logWrapper{log: &l}
	return logger
}

// Info logs an info message
func (z *logWrapper) Info() *zerolog.Event {
	return z.log.Info()
}

// Error logs an error message
func (z *logWrapper) Error() *zerolog.Event {
	return z.log.Error()
}
