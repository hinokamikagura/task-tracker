package command

import (
	"flag"
	"time"

	"github.com/hinokamikagura/task-tracker/schemas"
)

func MarkDoneCommand(args []string) {
	markDoneCmd := flag.NewFlagSet("update", flag.ExitOnError)
	id := markDoneCmd.Uint("taskId", 0, "Id of task you want to update")

	markDoneCmd.Parse(args)

	taskList, err := ReadTasks()
	if err != nil {
		logger.Errorf("Error occured while read file: %v", err)
		return
	}
	var updatedTask schemas.Task
	for i := 0; i < len(taskList); i++ {
		if taskList[i].Id == *id {
			taskList[i].Status = "done"
			taskList[i].UpdatedAt = time.Now()
			updatedTask = taskList[i]
			break
		}
	}

	if updatedTask.Title == "" {
		logger.Errorf("Not found task you want to update status")
		return
	}

	if err := WriteTasks(taskList); err != nil {
		logger.Errorf("Error occured while upate Task: %v", err)
		return
	}

	logger.Infof("Task marked done successfully:")
	logger.TaskLog(updatedTask)
}
