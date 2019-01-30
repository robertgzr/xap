package queue

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var AddCommand = cli.Command{
	Name:        "add",
	Category:    "queue",
	Usage:       "add track(s) to the queue",
	ArgsUsage:   "TRACK...",
	Description: "TRACK can be a file or URL or - to read the list of tracks/URLs from stdin",
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

		tracks := []string{ctx.Args().First()}
		tracks = append(tracks, ctx.Args().Tail()...)

		// check if we read from stdin
		if tracks[0] == "-" {
			tracks = []string{}
			in := bufio.NewScanner(os.Stdin)
			for in.Scan() {
				tracks = append(tracks, in.Text())
			}
			if err := in.Err(); err != nil {
				return fmt.Errorf("Error reading from stdin: %s", err)
			}
		}

		switch {
		case ctx.Bool("append-only"):
			return c.LoadAppend(tracks...)
		case ctx.Bool("replace"):
			if c.Paused() {
				defer c.Play()
			}
			return c.LoadReplace(tracks...)
		case ctx.Bool("next"):
			return c.LoadNext(tracks...)
		default:
			return c.LoadPlay(tracks...)
		}
	},
}
