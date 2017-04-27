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
		Name:  "xap",
		Usage: "cli for the xapper lib",
		Commands: []cli.Command{
			daemon.Command(),
			player.Command(),
			queue.Command(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "socket, s",
				Usage: "filepath to the ipc socket",
				Value: "/tmp/xapper.sock",
			},
		},
	}

	app.Run(os.Args)
}
