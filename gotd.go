package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/ghigt/gotd/task"
)

func runAction(c *cli.Context) {
	println("action run")
	println("time: ", c.Int("time"))
}

func listAction(c *cli.Context) {
	fmt.Printf("%v", tasks)
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
		fmt.Printf("added task: %q\n", name)
	}
}

func rmAction(c *cli.Context) {
	var id int64

	if len(c.Args()) > 0 {
		id, _ = strconv.ParseInt(c.Args().First(), 0, 4)
	} else {
		cli.ShowCommandHelp(c, "add")
	}
	tasks = tasks.Remove(int(id))
	//fmt.Printf("removed task: %q\n", name)
}

func editAction(c *cli.Context) {
	println("action edit")
}

func resetAction(c *cli.Context) {
	tasks = task.Tasks{}
}

var tasks task.Tasks

func readFile() {
	buf, err := ioutil.ReadFile(".datast")
	if err == nil {
		err = json.Unmarshal(buf, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func saveInFile() {
	b, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(".datast", b, 0640)
	if err != nil {
		log.Fatal(err)
	}
}

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
		{
			Name:   "reset",
			Usage:  "reset the task list",
			Action: resetAction,
		},
	}

	readFile()
	defer saveInFile()
	app.Run(os.Args)
}
