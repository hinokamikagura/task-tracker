package main

import (
	"os"

	"github.com/hinokamikagura/task-tracker/command"
	"github.com/hinokamikagura/task-tracker/config"
)

func main() {
	config.Init()
	logger := config.GetLogger("task-tracker")
	if len(os.Args) < 2 {
		logger.Info(`
		----- Task Tracker Cli -----
		available command: add, update, delete, mark-in-progress, mark-done
		list, list done, list todo, list in-progress
		----------------------------
		`)
		return
	}

	command.InitCommand()
	cmd := os.Args[1]
	switch cmd {
	case "add":
		req := command.TaskReq{
			Title:       os.Args[2],
			Description: os.Args[3],
		}
		addRes := command.AddTaskCommand(req)
		if addRes.Success == true {
			logger.Infof("%v", *addRes.Task)
		}
		// logger.Infof("This is logger", req)
	}
}
