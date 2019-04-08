package queue

import (
	"fmt"
	"html/template"
	"os"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

var Command = cli.Command{
	Name:    "queue",
	Aliases: []string{"q", "ls"},
	Usage:   "Show the queue",
	Subcommands: []cli.Command{
		AddCommand,
		NextCommand,
		PrevCommand,
		removeCommand,
		clearCommand,
		moveCommand,
		shuffleCommand,
		gotoCommand,
		saveCommand,
		loadCommand,
	},
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}

		queue, err := c.Queue()
		if err != nil {
			return err
		}

		if len(queue) == 0 {
			fmt.Println("Queue is empty")
			return nil
		}

		return renderQueue(queue)
	},
}

func renderQueue(queue mp.Queue) error {
	tmpl := `QUEUE:{{ range . }}
| {{ printf "%02d" .Index }}: {{ if .Current }}*{{ else }} {{ end }} {{ .Title }}{{ end }}
|
| {{ len . }} track(s)
`
	t := template.Must(template.New("queue").Parse(tmpl))
	return t.Execute(os.Stdout, queue)
}
