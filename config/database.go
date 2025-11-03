package config

import "os"

func InitializeDatabase() error {
	logger := NewLogger("json")
	jsonPath := "./db/task.json"

	_, err := os.Stat(jsonPath)
	if os.IsNotExist(err) {
		logger.Info("Json file doesn't exit. Creating")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create(jsonPath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
