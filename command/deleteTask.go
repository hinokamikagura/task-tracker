package command

import (
	"flag"
	"slices"

	"github.com/hinokamikagura/task-tracker/schemas"
)

func DeleteTaskCommand(args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Uint("taskId", 0, "Task Id which you want to delete.")
	deleteCmd.Parse(args)

	taskList, err := ReadTasks()
	if err != nil {
		logger.Errorf("Error occured while read file: %v", err)
		return
	}
	if len(taskList) == 0 {
		logger.Errorf("There is no task on the database.")
		return
	}
	taskCnt := len(taskList)
	var deleteTask schemas.Task
	for index, task := range taskList {
		if task.Id == *id {
			deleteTask = task
			taskList = slices.Delete(taskList, index, index+1)
			break
		}
	}

	if taskCnt == len(taskList) {
		logger.Errorf("Not found task which you want to delete.")
		return
	}

	if err := WriteTasks(taskList); err != nil {
		logger.Errorf("Error occured while deleting task.")
		return
	}

	logger.Infof("Deleted task: %v", deleteTask)
}
