package command

import (
	"fmt"

	"github.com/urfave/cli"
)

func VersionPrinter(ctx *cli.Context) {
	fmt.Fprintf(ctx.App.Writer, "%v version %v (%v)\n", ctx.App.Name, ctx.App.Version, ctx.App.Metadata["build_date"])
}

func init() {
	cli.VersionPrinter = VersionPrinter
}
