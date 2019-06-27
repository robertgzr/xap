package player

import (
	"github.com/robertgzr/xap/mp"
	"github.com/urfave/cli"
)

var showCommand = cli.Command{
	Name:  "show",
	Usage: "show the mpv UI (toggles visibility when already shown)",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		video, err := c.GetProperty("video")
		if err != nil {
			return err
		}
		if video == "false" {
			return c.SetProperty("video", 1)
		} else {
			return c.SetProperty("video", 0)
		}
	},
}
