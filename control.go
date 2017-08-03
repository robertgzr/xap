package main

import (
	"fmt"
	"html"
	"html/template"
	"os"
	"time"

	"github.com/blang/mpv"
	"github.com/urfave/cli"
)

func ControlCommands() []*cli.Command {
	return []*cli.Command{
		nowCmd(),
		playCmd(),
		pauseCmd(),
		stopCmd(),
		from0Cmd(),
	}
}

func nowCmd() *cli.Command {
	return &cli.Command{
		Name:     "now",
		Category: "control",
		Usage:    "show currently playing song",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			meta, err := c.Now()
			if err != nil {
				return err
			}

			tmpl := `CURRENT:
| {{ .Title | unescape }}
| {{with .Artist }}{{ . }}{{ end }}
| {{with .Album }}{{ . }} ({{ .Date }}) {{ .Nr }}{{ end }}
|
| {{ with .Pos }}{{ timefmt .Current }} / {{ timefmt .Len }} ({{ printf "%.2f%%" .CurrentPerc }}){{ end }}
`
			t := template.New("now")
			t.Funcs(template.FuncMap(map[string]interface{}{
				"timefmt": func(d time.Duration) string {
					return fmt.Sprintf("%02d:%02d:%02d", int(d.Hours())%24, int(d.Minutes())%60, int(d.Seconds())%60)
				},
				"unescape": html.UnescapeString,
			}))
			t = template.Must(t.Parse(tmpl))
			return t.Execute(os.Stdout, meta)
		},
	}
}

func playCmd() *cli.Command {
	return &cli.Command{
		Name:     "play",
		Category: "control",
		Usage:    "start playing the current file",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.Play()
		},
	}
}

func from0Cmd() *cli.Command {
	return &cli.Command{
		Name:     "from0",
		Category: "control",
		Usage:    "restart playback of the current track",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.Seek(0, mpv.SeekModeAbsolute)
		},
	}
}

func pauseCmd() *cli.Command {
	return &cli.Command{
		Name:     "pause",
		Category: "control",
		Usage:    "pause the current file",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.SetPause(true)
		},
	}
}

func stopCmd() *cli.Command {
	return &cli.Command{
		Name:     "stop",
		Category: "control",
		Usage:    "stop the current file",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.Stop()
		},
	}
}
