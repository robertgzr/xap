package queue

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var saveCommand = cli.Command{
	Name:      "save",
	Category:  "QUEUE",
	Usage:     "Save the current playlist to a file",
	ArgsUsage: "PATH",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format, f",
			Usage: "playlist format to save as",
			Value: "m3u",
		},
	},
	Action: func(ctx *cli.Context) error {
		// c := com.Connect()
		path := ctx.Args().First()
		if path == "" {
			return fmt.Errorf("Please specify a path.")
		}

		// tracks, err := c.List()
		// if err != nil {
		//	return err
		// }
		// return playlist.Save(path)
		return fmt.Errorf("not implemented yet")
	},
}

var loadCommand = cli.Command{
	Name:      "load",
	Category:  "QUEUE",
	Usage:     "Load playlist from a file",
	ArgsUsage: "PATH",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "append-only",
			Usage: "add but don't play",
		},
		&cli.BoolFlag{
			Name:  "replace, r",
			Usage: "add and play, stopping the current track",
		},
		&cli.BoolFlag{
			Name:  "next, n",
			Usage: "add as next track",
		},
	},
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}

		path := ctx.Args().First()
		if path == "" {
			return fmt.Errorf("Please specify a path.")
		}

		switch {
		case ctx.Bool("replace"):
			return c.LoadListReplace(path)
		case ctx.Bool("next"):
			return fmt.Errorf("not implemented yet")
			// return c.LoadNext(tracks...)
		default:
			return c.LoadListAppend(path)
		}
	},
}
