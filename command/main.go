package command

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/command/control"
	"github.com/robertgzr/xap/command/player"
	"github.com/robertgzr/xap/command/queue"
	"github.com/robertgzr/xap/command/raw"
)

func App(version, commit, date string) *cli.App {
	app := cli.NewApp()
	app.Name = "xap"
	app.Authors = []cli.Author{
		cli.Author{Name: "robertgzr", Email: "r@gnzler.io"},
	}
	app.Version = version
	app.Metadata = make(map[string]interface{})
	app.Metadata["build_commit"] = commit
	app.Metadata["build_date"] = date
	app.Usage = "cli to remote control mpv player"
	app.UsageText = "xap [global options] command [command options] [arguments...]"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "socket, S",
			Usage: "filepath to the ipc socket",
			Value: fmt.Sprintf("/var/run/user/%d/mpv", os.Getuid()),
		},
	}
	app.Commands = []cli.Command{
		control.Command,
		control.PlayCommand,
		control.PauseCommand,
		control.StopCommand,
		queue.Command,
		queue.NowPlayingCommand,
		queue.AddCommand,
		queue.NextCommand,
		queue.PrevCommand,
		player.Command,
		raw.Command,
	}
	app.Action = func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			cli.ShowAppHelpAndExit(ctx, 1)
		}
		return RunDynamicSubcommand(ctx)
	}
	return app
}
