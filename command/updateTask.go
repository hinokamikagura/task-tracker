package command

import (
	"flag"
	"time"

	"github.com/hinokamikagura/task-tracker/schemas"
)

func UpdateTaskCommand(args []string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	id := updateCmd.Uint("taskId", 0, "Id of task you want to update")
	title := updateCmd.String("title", "", "Title which you want to update")
	desc := updateCmd.String("desc", "", "Description which  you want to update")

	updateCmd.Parse(args)

	taskList, err := ReadTasks()
	if err != nil {
		logger.Errorf("Error occured while read file: %v", err)
		return
	}
	var updatedTask schemas.Task
	for i := 0; i < len(taskList); i++ {
		if taskList[i].Id == *id {
			if *title != "" {
				taskList[i].Title = *title
			}
			if *desc != "" {
				taskList[i].Description = *desc
			}
			taskList[i].UpdatedAt = time.Now()
			updatedTask = taskList[i]
			break
		}
	}

	if updatedTask.Title == "" {
		logger.Errorf("Not found task you want to update")
		return
	}

	if err := WriteTasks(taskList); err != nil {
		logger.Errorf("Error occured while upate Task: %v", err)
		return
	}

	logger.Infof("Updated task: %v", updatedTask)
}
