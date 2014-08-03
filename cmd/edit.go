package cmd

import (
	"log"

	"github.com/ghigt/cli"
)

var CmdEdit = cli.Command{
	Name:  "edit",
	Usage: "edit a task to the list",
	Flags: []cli.Flag{
		cli.StringFlag{"name, n", "", "new task name", "GOTD_N"},
	},
	Action: editAction,
}

func editAction(c *cli.Context) {
	var name, newName string

	if len(c.Args()) >= 2 {
		name = c.Args()[0]
		newName = c.Args()[1]
	} else {
		cli.ShowCommandHelp(c, "edit")
		return
	}
	task, err := Tasks.GetByName(name)
	if err != nil {
		log.Fatal(err)
	}
	task.Name = newName
}
