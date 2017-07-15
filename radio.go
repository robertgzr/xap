package main

import (
	"html/template"
	"os"

	"github.com/robertgzr/xap/pkg/radio"
	"github.com/urfave/cli"
)

func RadioCommand() *cli.Command {
	cmd := &cli.Command{
		Name:    "rad.io",
		Aliases: []string{"rad"},
		Usage:   "interface to r-a-d.io",
		Action:  radioStatus,
	}
	cmd.Subcommands = append(cmd.Subcommands, RadioPlayCommand())
	return cmd
}

func radioStatus(ctx *cli.Context) error {
	now, err := radio.Now()
	if err != nil {
		return err
	}

	tmpl := `r-a-d.io:
| {{ .NowPlaying }} {{ if .Pos.Ok }}| {{ .Pos.Current }} / {{ .Pos.Len }} ({{ printf "%.2f%%" .Pos.CurrentPerc }}){{ end }}
|
| {{ .Dj.Name }} {{ if .IsAfk }}(afk){{ end }}
| Listeners: {{ .Listeners }}
|
| Last Played:{{ range .LastPlayed }}
|   * {{ .Meta }}{{ end }}
| Queue:{{ range .Queue }}
|   * {{ .Meta }}{{ end }}
`
	// | Last: {{ with index .LastPlayed 0 }}{{ .Meta }}{{ end }}
	// | Next: {{ with index .Queue 0 }}{{ .Meta }}{{ end }}
	t := template.Must(template.New("now").Parse(tmpl))
	return t.Execute(os.Stdout, now)
}

func RadioPlayCommand() *cli.Command {
	return &cli.Command{
		Name:    "play",
		Aliases: []string{"p"},
		Usage:   "basically doing `xap add -r https://r-a-d.io/main.mp3`",
		Before:  initCom,
		Action: func(_ *cli.Context) error {
			return c.LoadReplace("https://r-a-d.io/main.mp3")
		},
	}
}
