package main

import (
	"os"

	"github.com/codegangsta/cli"
)

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
			Action: func(c *cli.Context) {
				println("action run")
				println("time: ", c.Int("time"))
			},
		},
		{
			Name:  "list",
			Usage: "list all tasks (id / title / due date)",
			Action: func(c *cli.Context) {
				println("action list")
			},
		},
		{
			Name:  "add",
			Usage: "add a task to the list",
			Action: func(c *cli.Context) {
				println("action add")
			},
		},
		{
			Name:  "rm",
			Usage: "remove a task to the list",
			Action: func(c *cli.Context) {
				println("action remove")
			},
		},
		{
			Name:  "edit",
			Usage: "edit a task to the list",
			Flags: []cli.Flag{
				cli.StringFlag{"name, n", "", "new task name"},
			},
			Action: func(c *cli.Context) {
				println("action edit")
			},
		},
	}

	app.Run(os.Args)
}
