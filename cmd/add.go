package cmd

import (
	"fmt"
	"time"

	"github.com/ghigt/cli"
)

var CmdAdd = cli.Command{
	Name:        "add",
	Usage:       "add a task to the list",
	Description: "Precise the `name` of the task you want to add.",
	Flags: []cli.Flag{
		cli.DurationFlag{"duration, d", time.Duration(25 * time.Minute), "add duration for the run"},
	},
	Action: addAction,
}

func addAction(c *cli.Context) {
	var name string

	if len(c.Args()) > 0 {
		name = c.Args().First()
	} else {
		cli.ShowCommandHelp(c, "add")
	}
	if len(name) > 0 {
		t := Tasks.Add(name, c.Duration("duration"))
		fmt.Printf("added task: %v\n", *t)
	}
}
