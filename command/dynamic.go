package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func RunDynamicSubcommand(ctx *cli.Context) error {
	cName := fmt.Sprintf("xap-%s", ctx.Args()[0])
	command, err := exec.LookPath(cName)
	if err != nil {
		return err
	}
	c := exec.Command(command, fmt.Sprintf("--socket=%s", ctx.GlobalString("socket")))
	c.Args = append(c.Args, ctx.Args()[1:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
