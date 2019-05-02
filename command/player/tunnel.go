package player

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var tunnelCommand = cli.Command{
	Name:  "tunnel",
	Usage: "tells xap to create a SSH based tunnel before to connect through, [user@]host:/path/to/socket",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "remote-socket, s",
			Usage: "location of the remote socket",
		},
		cli.StringFlag{
			Name:  "destination, H",
			Usage: "[user@]hostname[:port]",
		},
	},
	Action: func(ctx *cli.Context) error {
		localSocket := ctx.GlobalString("socket")
		remoteSocket := ctx.String("remote-socket")
		if remoteSocket == "" {
			return errors.New("remote-socket can't be empty")
		}
		addr := ctx.String("destination")
		if addr == "" {
			return errors.New("destination can't be empty")
		}

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGHUP)

		cmd := exec.Command("ssh", "-nNT", "-L", localSocket+":"+remoteSocket, addr)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			return errors.Wrapf(err, "failed to connect to %s", addr)
		}
		done := make(chan error, 1)
		go func() {
			err := cmd.Wait()
			done <- err
			close(done)
		}()
		fmt.Fprintf(os.Stdout, "Starting SSH tunnel... connect to %s\n", localSocket)

		select {
		case <-sigchan:
			fmt.Fprintf(os.Stdout, "Stopping tunnel...\n")
			if err := cmd.Process.Kill(); err != nil {
				return errors.Wrap(err, "failed to kill ssh command")
			}
		case err := <-done:
			fmt.Fprintf(os.Stdout, "Tunnel stopped working...\n")
			if err != nil {
				return errors.Wrap(err, "ssh command failed")
			}
		}

		return os.Remove(localSocket)
	},
}
