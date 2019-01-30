package queue

import (
	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var NextCommand = cli.Command{
	Name:     "next",
	Category: "queue",
	Usage:    "skips to the next track",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Next()
	},
}

var PrevCommand = cli.Command{
	Name:     "prev",
	Category: "queue",
	Usage:    "skips to the previous track",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Prev()
	},
}
