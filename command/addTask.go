package command

import (
	"encoding/json"
	"os"
	"time"

	"github.com/hinokamikagura/task-tracker/schemas"
)

type TaskReq struct {
	Title       string
	Description string
}

type AddTaskRes struct {
	Success bool          `json:"success"`
	Task    *schemas.Task `json:"task,omitempty"`
	Error   string        `json:"error,omitempty"`
}

func AddTaskCommand(task TaskReq) AddTaskRes {
	data, err := os.ReadFile("./db/task.json")
	if err != nil {
		logger.Errorf("Error reading file: %v", err)
		return AddTaskRes{Success: false, Error: err.Error()}
	}

	var taskList []schemas.Task

	if err := json.Unmarshal(data, &taskList); err != nil {
		logger.Errorf("Error parse json file: %v", err)
		return AddTaskRes{Success: false, Error: err.Error()}
	}

	newTask := schemas.Task{
		Id:          uint(len(taskList)),
		Title:       task.Title,
		Description: task.Description,
		Status:      "To Do",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskList = append(taskList, newTask)
	newData, err := json.MarshalIndent(taskList, "", "	")
	if err != nil {
		logger.Errorf("Error encoding JSON: %v", err)
		return AddTaskRes{Success: false, Error: err.Error()}
	}

	if err := os.WriteFile("./db/task.json", newData, 0644); err != nil {
		logger.Errorf("Error writting to JSON file: %v", err)
		return AddTaskRes{Success: false, Error: err.Error()}
	}

	logger.Infof("Task added successfully")
	return AddTaskRes{Success: true, Task: &newTask}
}
