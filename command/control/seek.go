package control

import (
	"strconv"

	"github.com/blang/mpv"
	"github.com/robertgzr/xap/mp"
	"github.com/urfave/cli"
)

var seekCommand = cli.Command{
	Name:            "seek",
	ArgsUsage:       "AMOUNT [MODE]",
	Category:        "CONTROL",
	Usage:           "Move playback to a different point in the current file",
	SkipFlagParsing: true,
	Action: func(ctx *cli.Context) error {
		var seekMode = mpv.SeekModeRelative
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}

		arg := ctx.Args().First()
		switch arg {
		case "-h":
			fallthrough
		case "--help":
			return cli.ShowCommandHelp(ctx, "seek")
		case "":
			return nil
		default:
			val, err := strconv.Atoi(arg)
			if err != nil {
				return err
			}
			if ctx.NArg() >= 2 {
				seekMode = ctx.Args().Get(1)
			}
			return c.Seek(val, seekMode)
		}
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
