package queue

import (
	"strconv"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var gotoCommand = cli.Command{
	Name:      "goto",
	Category:  "QUEUE",
	Usage:     "Start playing NUMBER track on the queue",
	ArgsUsage: "NUMBER",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}

		pos, err := strconv.Atoi(ctx.Args().First())
		if err != nil {
			return err
		}
		return c.Goto(pos)
	},
}
