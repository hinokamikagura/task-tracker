package config

import "fmt"

var (
	logger   *Logger
	filePath string
)

func Init() error {
	var err error
	filePath = "./db/task.json"
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

func GetFilePath() string {
	return filePath
}
