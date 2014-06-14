package cmd

import (
	"fmt"

	"github.com/ghigt/cli"
	"github.com/ghigt/gotd/task"
)

var CmdReset = cli.Command{
	Name:   "reset",
	Usage:  "reset the tasks list",
	Action: resetAction,
}

func resetAction(c *cli.Context) {
	Tasks = task.Tasks{}
	fmt.Println("tasks reseted")
}
