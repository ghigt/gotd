package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/ghigt/gotd/task"
	"github.com/ghigt/gotd/term"
)

func mySelect(t int) {
	tick := time.Tick(1 * time.Second)

	var current time.Duration
	var f time.Duration
	fmt.Println()
	for {
		select {
		case <-tick:
			current += 1 * time.Second
		default:
			if err := term.SetCap("up"); err != nil {
				log.Println(err)
			}
			if err := term.SetCap("ce"); err != nil {
				log.Println(err)
			}
			f = time.Duration(t)*time.Second - current
			fmt.Println(f)
			if f <= 0 {
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func runAction(c *cli.Context) {
	var name string

	if len(c.Args()) > 0 {
		name = c.Args().First()
	} else {
		cli.ShowCommandHelp(c, "run")
	}
	t, err := tasks.GetByName(name)
	if err != nil {
		log.Fatal(err)
	}
	if err := term.TGetEnt(); err != nil {
		log.Println(err)
	}
	mySelect(c.Int("time"))
	tasks.Remove(t.Id)
	fmt.Println("finished!")
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
		tasks.Add(name)
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
	task, err := tasks.Get(int(id))
	if err != nil {
		log.Fatal(err)
	}
	err = tasks.Remove(int(id))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("removed task: %q\n", task)
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
	task, err := tasks.GetByName(name)
	if err != nil {
		log.Fatal(err)
	}
	task.Name = newName
}

func resetAction(c *cli.Context) {
	tasks = task.Tasks{}
	fmt.Println("tasks reseted")
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
	app.Version = "0.1"
	app.Usage = "fight the laziness!"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Commands = []cli.Command{
		{
			Name: "run",
			//Usage: "randomly chose a task for the next given time",
			Usage: "choose a task for a given time",
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
