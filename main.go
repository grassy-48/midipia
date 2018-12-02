package main

import (
	"midipia/actions/input"
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
			Name:   "input-score",
			Usage:  "標準入力を受け付ける",
			Action: input.Input,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "出力するファイル名",
				},
			},
		},
		{
			Name:   "sample-read",
			Usage:  "midiを読み込む(sample)",
			Action: read.Read,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "入力するファイル名",
				},
			},
		},
		{
			Name:   "sample-create",
			Usage:  "midiを書き出す(sample)",
			Action: create.Create,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "出力するファイル名",
				},
			},
		},
	}
	app.Run(os.Args)
}
