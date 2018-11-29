package main

import (
	"midipia/actions/sample/create"
	"midipia/actions/sample/read"
	"os"
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	version = "0.0.1"
	appName = "midipia"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	app.Compiled = time.Now()

	app.Commands = []cli.Command{
		{
			Name:   "sample-read",
			Usage:  "midiを読み込む(sample)",
			Action: read.Read,
		},
		{
			Name:   "sample-create",
			Usage:  "midiを書き出す(sample)",
			Action: create.Create,
		},
	}
	app.Run(os.Args)
}
