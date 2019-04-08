package control

import (
	"github.com/blang/mpv"
	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var Command = cli.Command{
	Name:    "control",
	Aliases: []string{"c"},
	Usage:   "Control mpv via socket",
	Subcommands: []cli.Command{
		PlayCommand,
		PauseCommand,
		StopCommand,

		volumeCommand,
		from0Command,
	},
}

var PlayCommand = cli.Command{
	Name:     "play",
	Category: "CONTROL",
	Usage:    "Start to play the current file",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Play()
	},
}

var PauseCommand = cli.Command{
	Name:     "pause",
	Category: "CONTROL",
	Usage:    "Pause the current file",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.SetPause(true)
	},
}

var StopCommand = cli.Command{
	Name:     "stop",
	Category: "CONTROL",
	Usage:    "Stop the current file",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Stop()
	},
}

var from0Command = cli.Command{
	Name:     "from0",
	Category: "CONTROL",
	Usage:    "Restart playback of the current file",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Seek(0, mpv.SeekModeAbsolute)
	},
}
