package raw

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var Command = cli.Command{
	Name:      "raw",
	Usage:     "Raw communication interface",
	UsageText: `Refer to https://mpv.io/manual/master/#json-ipc`,
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
	Usage:     "Run raw mpv input commands",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		resp, err := c.Exec(toEmptyInterfaceSlice(ctx.Args())...)
		if err != nil {
			return err
		}
		if resp.Err != "success" {
			return errors.New(resp.Err)
		}
		if resp.Data != nil {
			return json.NewEncoder(os.Stdout).Encode(resp.Data)
		}
		fmt.Fprintf(os.Stdout, "%v\n", resp.Err)
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
	ArgsUsage: "PROPERTY",
	Usage:     "Read mpv properties",
	Description: `See mpv --list-properties for available PROPERTY values.
   Responses are formatted as JSON.

EXAMPLES:
   $ xap raw get mpv-version
`,
}

var setCommand = cli.Command{
	Name:      "set",
	ArgsUsage: "PROPERTY VALUE",
	Usage:     "Write mpv properties",
	Description: `See mpv --list-properties for available PROPERTY/VALUE values.
   Responses are formatted as JSON.

EXAMPLES:
   xap raw set paused 1
`,
}
