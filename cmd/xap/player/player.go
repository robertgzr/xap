package player

import (
	"fmt"

	"github.com/urfave/cli"
	"rbg.re/robertgzr/xapper/pkg/com"
)

var c *com.Com

func Command() cli.Command {
	cmd := cli.Command{
		Name:        "player",
		ShortName:   "p",
		Usage:       "control the mpv player instance",
		Description: "The subcommands in the playlist category are deprecated and are going to be replaced by `xapper queue` which will offer a much richer playback queue experience.",
		Subcommands: []cli.Command{
			playCmd(),
			pauseCmd(),
			stopCmd(),
			loadCmd(),
			nextCmd(),
			prevCmd(),
			showPlaylistCmd(),
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
		Name:     "play",
		Usage:    "start playing the current file",
		Category: "playback",
		Action: func(_ *cli.Context) error {
			return c.SetPause(false)
		},
	}
}

func pauseCmd() cli.Command {
	return cli.Command{
		Name:     "pause",
		Usage:    "pause the current file",
		Category: "playback",
		Action: func(_ *cli.Context) error {
			return c.SetPause(true)
		},
	}
}

func stopCmd() cli.Command {
	return cli.Command{
		Name:     "stop",
		Usage:    "stop the current file",
		Category: "playback",
		Action: func(_ *cli.Context) error {
			return c.Stop()
		},
	}
}

func loadCmd() cli.Command {
	return cli.Command{
		Name:      "load",
		Usage:     "load a track into mpv's internal playlist.",
		ArgsUsage: "FILE is the track to load",
		Category:  "playlist",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "mode, m",
				Usage: "play mode [append, append-play, replace]",
				Value: "append",
			},
		},
		Action: func(ctx *cli.Context) error {
			return c.Load(ctx.Args().First(), ctx.String("mode"))
		},
	}
}

func nextCmd() cli.Command {
	return cli.Command{
		Name:     "next",
		Usage:    "skips to the next track",
		Category: "playlist",
		Action: func(_ *cli.Context) error {
			return c.Next()
		},
	}
}

func prevCmd() cli.Command {
	return cli.Command{
		Name:     "prev",
		Usage:    "skips to the previous track",
		Category: "playlist",
		Action: func(_ *cli.Context) error {
			return c.Prev()
		},
	}
}

func showPlaylistCmd() cli.Command {
	return cli.Command{
		Name:      "playlist",
		ShortName: "ls",
		Usage:     "prints mpv's internal playlist",
		Category:  "playlist",
		Action: func(_ *cli.Context) error {
			ls, err := c.List()
			if err != nil {
				return err
			}
			fmt.Printf("PLAYLIST:\n%s\n", ls)
			return nil
		},
	}
}
