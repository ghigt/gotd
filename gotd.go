package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghigt/cli"
	"github.com/ghigt/gotd/cmd"
)

func readFile() {
	buf, err := ioutil.ReadFile(".datast")
	if err == nil {
		err = json.Unmarshal(buf, &cmd.Tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func saveInFile() {
	b, err := json.Marshal(cmd.Tasks)
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
		cmd.CmdRun,
		cmd.CmdList,
		cmd.CmdAdd,
		cmd.CmdRemove,
		cmd.CmdEdit,
		cmd.CmdReset,
	}

	readFile()
	defer saveInFile()
	app.Run(os.Args)
}
