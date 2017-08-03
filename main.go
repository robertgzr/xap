package main

import (
	"os"
	"time"

	"github.com/robertgzr/xap/pkg/com"

	"github.com/urfave/cli"
)

const version string = "v0.1.0"

var c *com.Com

func main() {
	app := &cli.App{
		Name:        "xap",
		Version:     "0.1.0",
		Compiled:    time.Now(),
		HideVersion: true,
		Usage:       "cli to remote control mpv player",
		UsageText:   "xap [global options] command [command options] [arguments...]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "socket",
				Aliases: []string{"S"},
				Usage:   "filepath to the ipc socket",
				Value:   "/tmp/xap.sock",
			},
		},
		EnableShellCompletion: true,
		// Action: func(c *cli.Context) error {
		// 	cli.DefaultAppComplete(c)
		// 	return nil
		// },
		// Authors: []cli.Author{
		//	cli.Author{Name: "robertgzr", Email: "robertguenzler@gmail.com"},
		// },
	}

	mountCommands(app, SettingsCommand())
	mountCommands(app, DaemonCommand())
	mountCommands(app, QueueCommands()...)
	mountCommands(app, ControlCommands()...)
	mountCommands(app, RawCommand())
	mountCommands(app, RadioCommand())

	app.Run(os.Args)
}

func initCom(ctx *cli.Context) (err error) {
	c, err = com.NewCom(ctx.String("socket"))
	return
}

func mountCommands(app *cli.App, subcmds ...*cli.Command) {
	app.Commands = append(app.Commands, subcmds...)
}
