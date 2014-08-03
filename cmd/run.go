package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/ghigt/cli"
	"github.com/ghigt/gotd/term"
)

var CmdRun = cli.Command{
	Name: "run",
	//Usage: "randomly chose a task for the next given time",
	Usage: "choose a task for the run",
	Flags: []cli.Flag{
		cli.DurationFlag{"duration, d", 0, "specify duration for the run.", "GOTD_D"},
	},
	Action: runAction,
}

func mySelect(t time.Duration) {
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
			f = t - current
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
	d := c.Duration("duration")

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
	if d == 0 {
		mySelect(t.Duration)
	} else {
		mySelect(d)
	}
	Tasks.Remove(t.Id)
	fmt.Println("finished!")
}
