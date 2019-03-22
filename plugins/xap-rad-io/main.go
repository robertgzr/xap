package main // import "github.com/robertgzr/xap/plugins/xap-rad-io"

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"

	xap "github.com/robertgzr/xap/command"
	"github.com/urfave/cli"
)

var (
	version   string
	buildInfo string
)

func init() {
	cli.VersionPrinter = xap.VersionPrinter
}

func main() {
	app := cli.NewApp()
	app.Name = "xap-rad-io"
	app.Usage = "interface to r-a-d.io"
	app.Version = version
	app.Metadata = make(map[string]interface{})
	app.Metadata["buildInfo"] = buildInfo
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "socket, S",
			Usage: "filepath to the ipc socket",
			Value: "/tmp/mpv.sock",
		},
	}
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
	Name:    "now-playing",
	Aliases: []string{"now", "status"},
	Action: func(ctx *cli.Context) error {
		now, err := NowPlaying()
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
	},
}

var playCommand = cli.Command{
	Name:    "play",
	Aliases: []string{"p"},
	Usage:   "basically doing `xap add -r https://r-a-d.io/main.mp3`",
	Action: func(ctx *cli.Context) error {
		c := exec.Command("xap", "--socket="+ctx.GlobalString("socket"), "add", "-r", StreamURL())
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		fmt.Println("+", c.Args)
		return c.Run()
	},
}
