package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ghigt/gotd/task"
)

func runAction(c *cli.Context) {
	println("action run")
	println("time: ", c.Int("time"))
}

func listAction(c *cli.Context) {
	println("action list")
}

func addAction(c *cli.Context) {
	var name string

	if len(c.Args()) > 0 {
		name = c.Args().First()
	} else {
		cli.ShowCommandHelp(c, "add")
	}
	if len(name) > 0 {
		tasks = tasks.Add(name)
		fmt.Println("added task:", name)
	}
}

func rmAction(c *cli.Context) {
	println("action remove")
}

func editAction(c *cli.Context) {
	println("action edit")
}

var tasks task.Tasks

func main() {
	app := cli.NewApp()
	app.Name = "gotd"
	app.Usage = "fight the lazyness!"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "randomly chose a task for the next given time",
			Flags: []cli.Flag{
				cli.IntFlag{"time, t", 25, "default time for the run"},
			},
			Action: runAction,
		},
		{
			Name:   "list",
			Usage:  "list all tasks (id / title / due date)",
			Action: listAction,
		},
		{
			Name:        "add",
			Usage:       "add a task to the list",
			Description: "Precise the `name` of the task you want to add.",
			Action:      addAction,
		},
		{
			Name:   "rm",
			Usage:  "remove a task to the list",
			Action: rmAction,
		},
		{
			Name:  "edit",
			Usage: "edit a task to the list",
			Flags: []cli.Flag{
				cli.StringFlag{"name, n", "", "new task name"},
			},
			Action: editAction,
		},
	}

	app.Run(os.Args)
}
