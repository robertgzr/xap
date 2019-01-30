package queue

import (
	"strconv"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var moveCommand = cli.Command{
	Name:      "move",
	Aliases:   []string{"mv"},
	Category:  "queue",
	Usage:     "moves a track from FROM to TO on the playlist",
	ArgsUsage: "FROM TO",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		from, err := strconv.Atoi(ctx.Args().Get(0))
		if err != nil {
			return err
		}
		to, err := strconv.Atoi(ctx.Args().Get(1))
		if err != nil {
			return err
		}
		return c.Move(from, to)
	},
}
