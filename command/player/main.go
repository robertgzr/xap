package player

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:    "player",
	Aliases: []string{"p"},
	Usage:   "Control a mpv process",
	Subcommands: []cli.Command{
		statusCommand,
		runCommand,
		stopCommand,
		showCommand,
		tunnelCommand,
	},
}

var statusCommand = cli.Command{
	Name:  "status",
	Usage: "inspect the status of a detached mpv command",
	Action: func(_ *cli.Context) error {
		pidStr, err := ioutil.ReadFile(pidFile)
		if err != nil {
			return err
		}
		pid, err := strconv.Atoi(string(pidStr))
		if err != nil {
			return err
		}
		p, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		cmdline, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", p.Pid))
		if err != nil {
			return err
		}
		args := bytes.Split(cmdline, []byte{0})
		if string(args[0]) != "mpv" {
			return errors.New("exited")
		}
		fmt.Fprintf(os.Stdout, "Runnning (pid: %d)\n", p.Pid)
		return nil
	},
}
