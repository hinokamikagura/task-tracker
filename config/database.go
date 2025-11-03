package config

import "os"

func InitializeDatabase() error {
	logger := GetLogger("json")

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		logger.Info("Json file doesn't exit. Creating")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}

		if err := os.WriteFile(filePath, []byte("[]"), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
