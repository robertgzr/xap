package main

import (
	"fmt"
	"os"

	"github.com/robertgzr/xap/command"
)

func main() {
	app := command.App()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", app.Name, err)
		os.Exit(1)
	}
}
