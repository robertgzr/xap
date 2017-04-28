package player

import (
	"github.com/urfave/cli"
	"rbg.re/robertgzr/xapper/pkg/com"
)

var c *com.Com

func Command() cli.Command {
	cmd := cli.Command{
		Name:      "player",
		ShortName: "p",
		Usage:     "control the playback of mpv",
		Subcommands: []cli.Command{
			playCmd(),
			pauseCmd(),
			stopCmd(),
		},
		Before: func(ctx *cli.Context) error {
			var err error
			c, err = com.NewCom(ctx.GlobalString("socket"))
			if err != nil {
				return err
			}
			return nil
		},
		Action: playerStatus,
	}

	return cmd
}

func playerStatus(ctx *cli.Context) error {
	// TODO: print player status
	println("not implemented")
	return nil
}

func playCmd() cli.Command {
	return cli.Command{
		Name:  "play",
		Usage: "start playing the current file",
		Action: func(_ *cli.Context) error {
			return c.Play()
		},
	}
}

func pauseCmd() cli.Command {
	return cli.Command{
		Name:  "pause",
		Usage: "pause the current file",
		Action: func(_ *cli.Context) error {
			return c.SetPause(true)
		},
	}
}

func stopCmd() cli.Command {
	return cli.Command{
		Name:  "stop",
		Usage: "stop the current file",
		Action: func(_ *cli.Context) error {
			return c.Stop()
		},
	}
}
