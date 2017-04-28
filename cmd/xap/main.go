package main

import (
	"os"

	"rbg.re/robertgzr/xapper/cmd/xap/daemon"
	"rbg.re/robertgzr/xapper/cmd/xap/player"
	"rbg.re/robertgzr/xapper/cmd/xap/queue"

	"github.com/urfave/cli"
)

const version string = "v0.1.0"

func main() {
	app := &cli.App{
		Name:      "xap",
		Usage:     "cli for the xapper lib",
		UsageText: "xap [global options] command [command options] [arguments...]",
		Commands: []cli.Command{
			daemon.Command(),
			player.Command(),
			queue.Command(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "socket, S",
				Usage: "filepath to the ipc socket",
				Value: "/tmp/xapper.sock",
			},
		},
		Version:              "0.1.0",
		HideVersion:          true,
		EnableBashCompletion: true,
		// Authors: []cli.Author{
		//	cli.Author{Name: "robertgzr", Email: "robertguenzler@gmail.com"},
		// },
	}

	app.Run(os.Args)
}
