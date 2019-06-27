package main

import (
	"fmt"
	"os"

	"github.com/robertgzr/xap/command"
)

var (
	version   = "undefined"
	buildInfo = "undefined"
)

func main() {
	app := command.App(version, buildInfo)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
		os.Exit(1)
	}
}
