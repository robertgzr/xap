package daemon

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
	"--no-softvol",
}

func Command() cli.Command {
	return cli.Command{
		Name:      "daemon",
		ShortName: "d",
		Usage:     "control the lifetime of a mpv process",
		Subcommands: []cli.Command{
			startCmd(),
			stopCmd(),
		},
		Action: daemonStatus,
	}
}

func daemonStatus(ctx *cli.Context) error {
	// TODO: show daemon status
	println("not implemented")
	return nil
}

func startCmd() cli.Command {
	return cli.Command{
		Name: "start",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-defaults",
				Usage: "do not run with default flags",
			},
		},
		Action: runDeamonStart,
	}
}

func runDeamonStart(ctx *cli.Context) error {
	var args = []string{
		"--idle",
		fmt.Sprintf("--input-ipc-server=%s", ctx.GlobalString("socket")),
	}
	if !ctx.Bool("no-defaults") {
		args = append(args, defaultMpvFlags...)
	}

	cmd := exec.Command("mpv", args...)
	fmt.Println("+", cmd.Args)

	if err := cmd.Start(); err != nil {
		return err
	}

	fmt.Printf("started: %d\n", cmd.Process.Pid)
	return nil
}

func stopCmd() cli.Command {
	return cli.Command{
		Name: "stop",
		Action: func(ctx *cli.Context) error {
			c, err := com.NewCom(ctx.GlobalString("socket"))
			if err != nil {
				return err
			}
			return c.Quit()
		},
	}
}
