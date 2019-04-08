package queue

import (
	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var shuffleCommand = cli.Command{
	Name:     "shuffle",
	Category: "QUEUE",
	Usage:    "Shuffle the current playlist",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		return c.Shuffle()
	},
}
