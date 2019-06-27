package main

import (
	"fmt"
	"os"

	"github.com/robertgzr/xap/command"
)

var (
	version = "undefined"
	commit  = "undefined"
	date    = "undefined"
)

func main() {
	app := command.App(version, commit, date)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
		os.Exit(1)
	}
}
