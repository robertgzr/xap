package raw

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var Command = cli.Command{
	Name:  "raw",
	Usage: "raw communication interface",
	Subcommands: []cli.Command{
		execCommand,
		getCommand,
		setCommand,
	},
}

var execCommand = cli.Command{
	Name:      "exec",
	Aliases:   []string{"x"},
	ArgsUsage: "COMMAND [ARGS...]",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		resp, err := c.Exec(toEmptyInterfaceSlice(ctx.Args())...)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%v", resp.Err)
		return nil
	},
}

func toEmptyInterfaceSlice(args cli.Args) []interface{} {
	var ifs = make([]interface{}, len(args))
	for i, a := range args {
		ifs[i] = a
	}
	return ifs
}

var getCommand = cli.Command{
	Name:      "get",
	ArgsUsage: "get PROPERTY",
	Usage:     "read mpv properties",
}

var setCommand = cli.Command{
	Name:      "set",
	ArgsUsage: "set PROPERTY VALUE",
	Usage:     "write mpv properties",
}
