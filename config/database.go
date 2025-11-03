package config

import "os"

func InitializeDatabase() error {
	logger := GetLogger("json")
	jsonPath := "./db/task.json"

	_, err := os.Stat(jsonPath)
	if os.IsNotExist(err) {
		logger.Info("Json file doesn't exit. Creating")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}

		if err := os.WriteFile(jsonPath, []byte("[]"), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
