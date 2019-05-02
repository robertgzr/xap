package main // import "github.com/robertgzr/xap/plugins/xap-radio"

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"

	xap "github.com/robertgzr/xap/command"
	"github.com/robertgzr/xap/mp"
	"github.com/robertgzr/xap/plugins/xap-radio/radio"
	"github.com/robertgzr/xap/plugins/xap-radio/stew"
)

var (
	version   = "undefined"
	buildInfo = "undefined"
)

func init() {
	cli.VersionPrinter = xap.VersionPrinter
}

func main() {
	app := cli.NewApp()
	app.Name = "xap-radio"
	app.Usage = "interface to r-a-d.io and stew.moe"
	app.Version = version
	app.Metadata = make(map[string]interface{})
	app.Metadata["buildInfo"] = buildInfo
	app.Flags = xap.App("", "").Flags
	app.Commands = []cli.Command{
		statusCommand,
		playCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
		os.Exit(1)
	}
}

var statusCommand = cli.Command{
	Name:    "playing",
	Aliases: []string{"now", "status"},
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		url, err := c.GetProperty("stream-open-filename")
		if err != nil {
			return err
		}
		switch strings.Trim(url, "\"") {
		case radio.StreamURL():
			radio.PrintStatus()
		case stew.StreamURL():
			stew.PrintStatus()
		default:
			fmt.Printf("unknown stream: %s\n", url)
		}
		return nil
	},
}

var playCommand = cli.Command{
	Name:    "play",
	Aliases: []string{"p"},
	Usage:   "basically doing `xap add -r <radio-stream-url>`",
	Description: `
   Supported web radio streams:
   - r-a-d.io
   - stew.moe
`,
	Action: func(ctx *cli.Context) error {
		var streamURL string

		switch ctx.Args().First() {
		case "r-a-d.io":
			streamURL = radio.StreamURL()
		case "stew.moe":
			streamURL = stew.StreamURL()
		default:
			if err := cli.ShowCommandHelp(ctx, ctx.Command.Name); err != nil {
				return err
			}
			return errors.New("unknown radio station")
		}

		c := exec.Command("xap", "--socket="+ctx.GlobalString("socket"), "add", "-r", streamURL)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		fmt.Println("+", c.Args)
		return c.Run()
	},
}
