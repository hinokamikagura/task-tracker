package command

import (
	"encoding/json"
	"os"

	"github.com/hinokamikagura/task-tracker/schemas"
)

func ReadTasks(filePath string) ([]schemas.Task, error) {
	data, err := os.ReadFile("./db/task.json")
	if err != nil {
		logger.Errorf("Error reading file: %v", err)
		return nil, err
	}

	var taskList []schemas.Task

	if err := json.Unmarshal(data, &taskList); err != nil {
		logger.Errorf("Error parse json file: %v", err)
		return nil, err
	}

	return taskList, nil
}

func WriteTasks(filePath string, tasks []schemas.Task) error {
	newData, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		logger.Errorf("Error encoding JSON: %v", err)
		return err
	}

	return os.WriteFile(filePath, newData, 0644)
}
