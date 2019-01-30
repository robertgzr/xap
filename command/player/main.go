package player

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var Command = cli.Command{
	Name:    "player",
	Aliases: []string{"p"},
	Usage:   "control a mpv process",
	Subcommands: []cli.Command{
		statusCommand,
		runCommand,
		stopCommand,
	},
}

var statusCommand = cli.Command{
	Name:  "status",
	Usage: "inspect the status of a detached mpv command",
}

var defaultMpvFlags = []string{
	"--idle",
	"--no-config",
	"--no-video",
	"--no-sub",
	"--gapless-audio=yes",
	"--volume=70.0",
}
var runCommand = cli.Command{
	Name:        "run",
	Usage:       "start an instance of mpv.",
	Description: `Everything after "--" is passed as arguments to mpv.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "detached, d",
			Usage: "run in the background",
		},
		&cli.BoolFlag{
			Name:  "no-defaults",
			Usage: "do not run with default flags",
		},
		&cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "print the mpv command",
		},
	},
	Action: func(ctx *cli.Context) error {
		args := append(ctx.Args(), fmt.Sprintf("--input-ipc-server=%s", ctx.GlobalString("socket")))
		if !ctx.Bool("no-defaults") {
			args = append(args, defaultMpvFlags...)
		}

		cmd := exec.Command("mpv", args...)
		if ctx.Bool("verbose") {
			fmt.Println("+", cmd.Args)
		}

		if ctx.Bool("detached") {
			if err := cmd.Start(); err != nil {
				return err
			}
			fmt.Fprintf(os.Stdout, "Started mpv... (pid: %d)\n", cmd.Process.Pid)
			return nil
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	},
}

var stopCommand = cli.Command{
	Name:  "stop",
	Usage: "stops the mpv instance",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		if err := c.Quit(); err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Stopped mpv...\n")
		return nil
	},
}
