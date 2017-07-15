package main

import (
	"html/template"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func SettingsCommand() *cli.Command {
	return &cli.Command{
		Name:  "settings",
		Usage: "configure mpv player options",
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
		Before:    initCom,
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
	tmpl := `AUDIO DEVICES:{{ range . }}
| {{ printf "%02d" .ID }}: {{ if .Current }}*{{ else }} {{ end }} {{ .Name }}{{ end }}
|
| {{ len . }} device(s)
`
	t := template.Must(template.New("ad").Parse(tmpl))
	return t.Execute(os.Stdout, ls)
}
