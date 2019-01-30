package queue

import (
	"strconv"
	"strings"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var removeCommand = cli.Command{
	Name:        "remove",
	Category:    "queue",
	Aliases:     []string{"rm"},
	Usage:       "remove tracks from the playlist",
	ArgsUsage:   "POSITION",
	Description: "POSITION can be a single index or a range expression like from..to (`to` is not removed)",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		ns := strings.Split(ctx.Args().First(), "..")
		switch {
		// POSITION is range expression
		case len(ns) == 2:
			from, err := strconv.Atoi(ns[0])
			if err != nil {
				return err
			}
			to, err := strconv.Atoi(ns[1])
			if err != nil {
				return err
			}
			for i := 0; i < (to - from); i++ {
				if err := c.Remove(from); err != nil {
					return err
				}
			}
			return nil
		// POSITION is single index
		case len(ns) == 1:
			n, err := strconv.Atoi(ns[0])
			if err != nil {
				return err
			}
			return c.Remove(n)
		// POSITION is invalid
		default:
			return nil
		}
	},
}

var clearCommand = cli.Command{
	Name:     "clear",
	Category: "queue",
	Usage:    "remove all entries from the queue",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Clear()
	},
}
