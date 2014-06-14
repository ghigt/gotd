package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ghigt/cli"
)

var CmdRemove = cli.Command{
	Name:   "rm",
	Usage:  "remove a task to the list",
	Action: rmAction,
}

func rmAction(c *cli.Context) {
	var id int64

	if len(c.Args()) > 0 {
		id, _ = strconv.ParseInt(c.Args().First(), 0, 4)
	} else {
		cli.ShowCommandHelp(c, "add")
	}
	Task, err := Tasks.Get(int(id))
	if err != nil {
		log.Fatal(err)
	}
	err = Tasks.Remove(int(id))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("removed task: %v\n", Task)
}
