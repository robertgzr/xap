package main

import (
	"fmt"

	"github.com/blang/mpv"
	"github.com/urfave/cli"
)

func RawCommand() *cli.Command {
	return &cli.Command{
		Name:      "raw",
		ArgsUsage: "[get|set] PROPERTY VALUE",
		Usage:     "manipulate a raw mpv property",
		Before:    initCom,
		Action: func(ctx *cli.Context) error {
			switch ctx.Args().First() {
			case "get":
				return getCmd(ctx)
			case "set":
				return setCmd(ctx)
			case "cycle":
				return cycleCmd(ctx)
			default:
				fmt.Printf("unknown operation\n")
				return nil
			}
		},
	}
}

func getCmd(ctx *cli.Context) error {
	prop := ctx.Args().Get(1)
	if prop == "" {
		return fmt.Errorf("no property given")
	}
	val, err := c.GetProperty(prop)
	if err != nil {
		return err
	}
	fmt.Println(val)
	return nil
}

func setCmd(ctx *cli.Context) error {
	prop := ctx.Args().Get(1)
	if prop == "" {
		return fmt.Errorf("no property given")
	}
	value := ctx.Args().Get(2)
	if value == "" {
		return fmt.Errorf("no value given")
	}
	return c.SetProperty(prop, value)
}

func cycleCmd(ctx *cli.Context) error {
	var res *mpv.Response
	var err error
	if direc := ctx.Args().Get(1); direc != "" {
		res, err = c.Exec("cycle", ctx.Args().Get(0), "up")
	}
	res, err = c.Exec("cycle", ctx.Args().Get(0), ctx.Args().Get(1))
	if err != nil {
		return fmt.Errorf("%s: %s", err, res.Err)
	}
	fmt.Printf("%+v", res.Data)
	return nil
}
