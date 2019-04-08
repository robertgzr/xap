package control

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var volumeCommand = cli.Command{
	Name:            "vol",
	ArgsUsage:       "[[+|-]VALUE]",
	Category:        "CONTROL",
	Usage:           "Print and adjust the softvol property",
	SkipFlagParsing: true,
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}

		arg := ctx.Args().First()
		switch arg {
		case "":
			vol, err := c.Volume()
			if err != nil {
				return err
			}
			fmt.Printf("VOLUME:\n| %v\n", vol)
			return nil
		case "-h":
			fallthrough
		case "--help":
			return cli.ShowCommandHelp(ctx, "vol")
		default:
			val, err := strconv.ParseFloat(arg[1:], 64)
			if err != nil {
				return err
			}
			switch arg[:1] {
			case "+":
				return c.VolumeUp(val)
			case "-":
				return c.VolumeDown(val)
			default:
				return fmt.Errorf("missing + or - before VALUE")
			}
		}
	},
}
