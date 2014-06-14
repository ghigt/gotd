package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/codegangsta/cli"
	"github.com/ghigt/gotd/term"
)

var CmdRun = cli.Command{
	Name: "run",
	//Usage: "randomly chose a task for the next given time",
	Usage: "choose a task for a given time",
	Flags: []cli.Flag{
		cli.IntFlag{"time, t", 25, "default time for the run"},
	},
	Action: runAction,
}

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
	t, err := Tasks.GetByName(name)
	if err != nil {
		log.Fatal(err)
	}
	if err := term.TGetEnt(); err != nil {
		log.Println(err)
	}
	mySelect(c.Int("time"))
	Tasks.Remove(t.Id)
	fmt.Println("finished!")
}
