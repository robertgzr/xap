package queue

import (
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
			Usage: "keep open and refresh",
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
		return renderNowPlaying(meta)
	},
}

func renderNowPlaying(meta mp.Metadata) error {
	tmpl := `NOW:
| {{ .Title | unescape }}
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
	return t.Execute(os.Stdout, meta)
}
