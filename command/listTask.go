package command

import "github.com/hinokamikagura/task-tracker/schemas"

func ListTaskCommand(status string) {
	taskList, err := ReadTasks()
	if err != nil {
		logger.Errorf("Error occured while read file: %v", err)
		return
	}

	if status == "all" {
		logger.TaskListLog(taskList)
	} else if status == "todo" || status == "done" || status == "in-progress" {
		var taskListoutput []schemas.Task
		for _, task := range taskList {
			if task.Status == status {
				taskListoutput = append(taskListoutput, task)
			}
		}
		logger.TaskListLog(taskListoutput)
	} else {
		logger.Errorf("Please input correct status!!!")
	}
}
