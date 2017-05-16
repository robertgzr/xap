package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

func PlayerSubcommand() *cli.Command {
	return &cli.Command{
		Name:    "player",
		Aliases: []string{"p"},
		Usage:   "configure mpv player options",
		Subcommands: []*cli.Command{
			listAudioDevicesCmd(),
			setAudioDeviceCmd(),
		},
		Before: initCom,
	}
}

func listAudioDevicesCmd() *cli.Command {
	return &cli.Command{
		Name:  "audio-device-list",
		Usage: "print available audio devices",
		Action: func(_ *cli.Context) error {
			ls, err := c.AudioDeviceList()
			if err != nil {
				return err
			}
			var buf bytes.Buffer
			buf.WriteString("AUDIO DEVICES" + ":\n")
			for _, d := range ls {
				var current string
				if d.Current {
					current = ">"
				}
				buf.WriteString(fmt.Sprintf("%2s %2d: %s\n", current, d.ID, d.Name))
			}
			buf.WriteString(fmt.Sprintf("\n%d device(s)", len(ls)))
			fmt.Println(buf.String())
			return nil
		},
	}

}

func setAudioDeviceCmd() *cli.Command {
	return &cli.Command{
		Name:      "audio-device",
		ArgsUsage: "DEVICE",
		Usage:     "change the audio device that mpv uses",
		Action: func(ctx *cli.Context) error {
			nr, err := strconv.Atoi(ctx.Args().First())
			if err != nil {
				return err
			}
			return c.SetAudioDevice(nr)
		},
	}
}
