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
		command.AddTaskCommand(os.Args[2:])
	case "delete":
		command.DeleteTaskCommand(os.Args[2:])
	case "update":
		command.UpdateTaskCommand(os.Args[2:])
	case "mark-in-progress":
		command.MarkProgressCommand(os.Args[2:])
	case "mark-done":
		command.MarkDoneCommand(os.Args[2:])
	}
}
