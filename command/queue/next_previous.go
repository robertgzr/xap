package queue

import (
	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var NextCommand = cli.Command{
	Name:     "next",
	Category: "QUEUE",
	Usage:    "Skips to the next title",
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
	Category: "QUEUE",
	Usage:    "Skips to the previous title",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Prev()
	},
}
