package queue

import (
	"encoding/json"
	"fmt"
	"html"
	"os"
	"text/template"
	"time"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var NowPlayingCommand = cli.Command{
	Name:     "now",
	Aliases:  []string{"status"},
	Category: "QUEUE",
	Usage:    "Show currently playing song",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "watch, w",
			Usage: "(not done yet) keep open and refresh",
		},
		&cli.BoolFlag{
			Name:  "json, j",
			Usage: "output json",
		},
	},
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		meta, err := c.Now()
		if err != nil {
			return err
		}
		if ctx.Bool("json") {
			return jsonNowPlaying(meta)
		}
		return renderNowPlaying(meta, c.Paused())
	},
}

func renderNowPlaying(meta mp.Metadata, paused bool) error {
	tmpl := `| {{ .Title | unescape }}
{{- with .Artist }}| {{ . }}{{ end }}
{{- with .Album }}| {{ . }} ({{ .Date }}) {{ .Nr }}{{ end }}
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
	fmt.Fprintf(os.Stdout, "NOW: %s\n", func() string {
		if paused {
			return "(paused)"
		} else {
			return ""
		}
	}())
	return t.Execute(os.Stdout, meta)
}

func jsonNowPlaying(meta mp.Metadata) error {
	return json.NewEncoder(os.Stdout).Encode(meta)
}
