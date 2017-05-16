package main

import (
	"fmt"
	"os/exec"

	"rbg.re/robertgzr/xapper/pkg/com"

	"github.com/urfave/cli"
)

var defaultMpvFlags = []string{
	"--no-config",
	"--no-video",
	"--no-sub",
	"--gapless-audio=yes",
	"--volume=50.0",
}

// DaemonSubcommand ...
func DaemonCommand() *cli.Command {
	return &cli.Command{
		Name:    "daemon",
		Usage:   "control the default mpv process",
		Subcommands: []*cli.Command{
			runDaemonCmd(),
			stopDaemonCmd(),
		},
		Action: daemonStatus,
	}
}

func daemonStatus(ctx *cli.Context) error {
	// TODO: show daemon status
	println("not implemented")
	return nil
}

func runDaemonCmd() *cli.Command {
	return &cli.Command{
		Name: "start",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "no-defaults",
				Usage: "do not run with default flags",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Aliases: []string{"v"},
				Usage: "print the mpv command",
			},
		},
		Action: func(ctx *cli.Context) error {
			var args = []string{
				"--idle",
				fmt.Sprintf("--input-ipc-server=%s", ctx.String("socket")),
			}
			if !ctx.Bool("no-defaults") {
				args = append(args, defaultMpvFlags...)
			}

			cmd := exec.Command("mpv", args...)

			if ctx.Bool("verbose") {
				fmt.Println("+", cmd.Args)
			}

			if err := cmd.Start(); err != nil {
				return err
			}
			fmt.Printf("started... (pid: %d)\n", cmd.Process.Pid)
			return nil
		},
	}
}

func stopDaemonCmd() *cli.Command {
	return &cli.Command{
		Name:   "stop",
		Before: initCom,
		Action: func(ctx *cli.Context) error {
			c, err := com.NewCom(ctx.String("socket"))
			if err != nil {
				return err
			}
			return c.Quit()
		},
	}
}
