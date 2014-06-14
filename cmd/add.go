package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var CmdAdd = cli.Command{
	Name:        "add",
	Usage:       "add a task to the list",
	Description: "Precise the `name` of the task you want to add.",
	Action:      addAction,
}

func addAction(c *cli.Context) {
	var name string

	if len(c.Args()) > 0 {
		name = c.Args().First()
	} else {
		cli.ShowCommandHelp(c, "add")
	}
	if len(name) > 0 {
		Tasks.Add(name)
		fmt.Printf("added task: %q\n", name)
	}
}
