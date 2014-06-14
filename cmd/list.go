package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var CmdList = cli.Command{
	Name:   "list",
	Usage:  "list all tasks (id / title / due date)",
	Action: listAction,
}

func listAction(c *cli.Context) {
	fmt.Printf("%v", Tasks)
}
