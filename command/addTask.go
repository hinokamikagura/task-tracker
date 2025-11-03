package command

import (
	"flag"
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

func AddTaskCommand(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	title := addCmd.String("title", "", "Title of the task (required)")
	desc := addCmd.String("desc", "", "Desction of the task (required)")

	addCmd.Parse(args)
	taskList, err := ReadTasks()
	if err != nil {
		logger.Errorf("Error occured while read file: %v", err)
		return
	}

	newTask := schemas.Task{
		Id:          uint(len(taskList)),
		Title:       *title,
		Description: *desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskList = append(taskList, newTask)
	if err := WriteTasks(taskList); err != nil {
		logger.Errorf("Error occured while write file: %v", err)
		return
	}
	logger.Infof("Successfully Added Task :")
	logger.TaskLog(newTask)
}
