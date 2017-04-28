package queue

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
	"rbg.re/robertgzr/xapper/pkg/com"
)

var c *com.Com

func Command() cli.Command {
	return cli.Command{
		Name:      "queue",
		ShortName: "q",
		Usage:     "manage mpv's internal playlist",
		Subcommands: []cli.Command{
			addCmd(),
			nextCmd(),
			prevCmd(),
			rmCmd(),
			clearCmd(),
			moveCmd(),
			shuffleCmd(),
		},
		Before: func(ctx *cli.Context) error {
			var err error
			c, err = com.NewCom(ctx.GlobalString("socket"))
			if err != nil {
				return err
			}
			return nil
		},
		Action: queueStatus,
	}
}

func queueStatus(ctx *cli.Context) error {
	ls, err := c.List()
	if err != nil {
		return err
	}

	if len(ls) == 0 {
		fmt.Println("Queue is empty")
		return nil
	}

	var buf bytes.Buffer
	buf.WriteString("QUEUE:\n")
	for _, tr := range ls {
		var current string
		if tr.Current {
			current = ">"
		}

		buf.WriteString(fmt.Sprintf("%2s %d: %s\n", current, tr.Index, tr.Title))
	}
	buf.WriteString(fmt.Sprintf("\n%d track(s)", len(ls)))
	fmt.Println(buf.String())
	return nil
}

func addCmd() cli.Command {
	return cli.Command{
		Name:        "add",
		Usage:       "add track(s) to the queue",
		ArgsUsage:   "TRACK...",
		Description: "TRACK can be a file or URL",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "append-only",
				Usage: "add but don't play",
			},
			cli.BoolFlag{
				Name:  "replace, r",
				Usage: "add and play, stopping the current track",
			},
			cli.BoolFlag{
				Name:  "next, n",
				Usage: "add as next track",
			},
		},
		Action: func(ctx *cli.Context) error {
			tracks := []string{ctx.Args().First()}
			tracks = append(tracks, ctx.Args().Tail()...)

			switch {
			case ctx.Bool("append-only"):
				return c.LoadAppend(tracks...)
			case ctx.Bool("replace"):
				return c.LoadReplace(tracks...)
			case ctx.Bool("next"):
				return c.LoadNext(tracks...)
			default:
				return c.LoadPlay(tracks...)
			}
		},
	}
}

func rmCmd() cli.Command {
	return cli.Command{
		Name:      "remove",
		ShortName: "rm",
		Usage:     "remove the track at position from the playlist",
		ArgsUsage: "POSITION",
		Action: func(ctx *cli.Context) error {
			n, err := strconv.Atoi(ctx.Args().First())
			if err != nil {
				return err
			}
			return c.Remove(n)
		},
	}
}

func clearCmd() cli.Command {
	return cli.Command{
		Name:  "clear",
		Usage: "remove all entries from the queue",
		Action: func(_ *cli.Context) error {
			return c.Clear()
		},
	}
}

func moveCmd() cli.Command {
	return cli.Command{
		Name:      "move",
		ShortName: "mv",
		Usage:     "moves a track from FROM to TO on the playlist",
		ArgsUsage: "FROM TO",
		Action: func(ctx *cli.Context) error {
			from, err := strconv.Atoi(ctx.Args().Get(0))
			if err != nil {
				return err
			}
			to, err := strconv.Atoi(ctx.Args().Get(1))
			if err != nil {
				return err
			}
			return c.Move(from, to)
		},
	}
}

func shuffleCmd() cli.Command {
	return cli.Command{
		Name:  "shuffle",
		Usage: "shuffle the current playlist",
		Action: func(_ *cli.Context) error {
			return c.Shuffle()
		},
	}
}

func nextCmd() cli.Command {
	return cli.Command{
		Name:  "next",
		Usage: "skips to the next track",
		Action: func(_ *cli.Context) error {
			return c.Next()
		},
	}
}

func prevCmd() cli.Command {
	return cli.Command{
		Name:  "prev",
		Usage: "skips to the previous track",
		Action: func(_ *cli.Context) error {
			return c.Prev()
		},
	}
}
