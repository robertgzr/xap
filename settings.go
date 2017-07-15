package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

func SettingsCommand() *cli.Command {
	return &cli.Command{
		Name:    "settings",
		Usage:   "configure mpv player options",
		Subcommands: []*cli.Command{
			audioDevicesCmd(),
		},
		Before: initCom,
	}
}

func audioDevicesCmd() *cli.Command {
	return &cli.Command{
		Name:      "audio-device",
		ArgsUsage: "DEVICE",
		Usage:     "change the audio device that mpv uses",
		Before: initCom,
		Action: func(ctx *cli.Context) error {
			device := ctx.Args().First()
			if device == "" {
				return listAudioDevices()
			}
			return setAudioDevice(device)
		},
	}

}

func setAudioDevice(device string) error {
	nr, err := strconv.Atoi(device)
	if err != nil {
		return err
	}
	return c.SetAudioDevice(nr)
}

func listAudioDevices() error {
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
}
