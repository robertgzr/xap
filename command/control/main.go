package control

import (
	"github.com/blang/mpv"
	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var Command = cli.Command{
	Name:    "control",
	Aliases: []string{"c"},
	Usage:   "control mpv via socket",
	Subcommands: []cli.Command{
		PlayCommand,
		PauseCommand,
		StopCommand,
		NowPlayingCommand,

		volumeCommand,
		from0Command,
	},
}

var PlayCommand = cli.Command{
	Name:  "play",
	Usage: "start playing the current file",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Play()
	},
}

var PauseCommand = cli.Command{
	Name:  "pause",
	Usage: "pause the current file",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.SetPause(true)
	},
}

var StopCommand = cli.Command{
	Name:  "stop",
	Usage: "stop the current file",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Stop()
	},
}

var from0Command = cli.Command{
	Name:  "from0",
	Usage: "restart playback of the current track",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Seek(0, mpv.SeekModeAbsolute)
	},
}
