package config

import "fmt"

var (
	logger *Logger
)

func Init() error {
	var err error

	err = InitializeDatabase()

	if err != nil {
		return fmt.Errorf("Error initializing json: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
