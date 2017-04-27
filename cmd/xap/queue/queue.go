package queue

import "github.com/urfave/cli"

func Command() cli.Command {
	return cli.Command{
		Name:      "queue",
		ShortName: "q",
		Action:    queueStatus,
	}
}

func queueStatus(ctx *cli.Context) error {
	// TODO: print the status of the queue
	println("not implemented")
	return nil
}
