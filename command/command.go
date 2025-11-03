package command

import (
	"github.com/hinokamikagura/task-tracker/config"
)

var (
	logger *config.Logger
)

func InitCommand() {
	logger = config.GetLogger("command")
}
