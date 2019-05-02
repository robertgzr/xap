package player

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/robertgzr/xap/mp"
	"github.com/urfave/cli"
)

var pidFile = fmt.Sprintf("/var/run/user/%d/xap-player.pid", os.Getuid())

var defaultMpvFlags = []string{
	"--no-video",
	"--no-sub",
	"--gapless-audio=yes",
	"--volume=70.0",
}
var runCommand = cli.Command{
	Name:        "run",
	Aliases:     []string{"start"},
	Usage:       "Start an instance of mpv.",
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
		args := append(ctx.Args(),
			fmt.Sprintf("--input-ipc-server=%s", ctx.GlobalString("socket")),
			"--idle",
		)
		if !ctx.Bool("no-defaults") {
			args = append(args, defaultMpvFlags...)
		}

		cmd := exec.Command("mpv", args...)
		if ctx.Bool("verbose") {
			fmt.Println("+", cmd.Args)
		}

		if !ctx.Bool("detached") {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}

		err := cmd.Start()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%v", cmd.Process.Pid)), 0666)
		if err != nil {
			if os.IsExist(err) {
				fmt.Fprintf(os.Stdout, "Already started\n")
				return cmd.Process.Kill()
			}
			return err
		}

		fmt.Fprintf(os.Stdout, "Started mpv (pid: %d)\n", cmd.Process.Pid)
		return nil
	},
}

var stopCommand = cli.Command{
	Name:  "stop",
	Usage: "Stops the mpv instance",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		if err := c.Quit(); err != nil {
			return err
		}
		if err := os.Remove(pidFile); err != nil {
			if os.IsNotExist(err) {
				return errors.New("No player running")
			}
			return err
		}
		fmt.Fprintf(os.Stdout, "Stopped mpv\n")
		return nil
	},
}
