package control

import (
	"encoding/json"
	"os"

	"github.com/urfave/cli"

	"github.com/robertgzr/xap/mp"
)

type status struct {
	Metadata mp.Metadata `json:"metadata"`
	Queue    mp.Queue    `json:"queue"`
	Paused   bool        `json:"paused"`
}

var statusCommand = cli.Command{
	Name:     "status",
	Category: "CONTROL",
	Usage:    "Get the current status of the player",
	Action: func(ctx *cli.Context) error {
		c, err := mp.Connect(ctx)
		if err != nil {
			return err
		}
		var r status
		r.Paused = c.Paused()
		r.Metadata, err = c.Now()
		if err != nil {
			return err
		}
		r.Queue, err = c.Queue()
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(r)
	},
}
