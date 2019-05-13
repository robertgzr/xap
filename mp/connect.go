package mp

import (
	"github.com/urfave/cli"
)

// TODO: remove this compat package and put stuff from pkg/com in here
func Connect(ctx *cli.Context) (*Mp, error) {
	client, err := NewMp(ctx.GlobalString("socket"))
	if err != nil {
		return nil, err
	}

	return client, nil
}
