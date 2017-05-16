package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

func QueueSubcommand() *cli.Command {
	return &cli.Command{
		Name:    "queue",
		Aliases: []string{"q"},
		Usage:   "manage mpv's internal playlist",
		Subcommands: []*cli.Command{
			addCmd(),
			nextCmd(),
			prevCmd(),
			rmCmd(),
			clearCmd(),
			moveCmd(),
			shuffleCmd(),
			gotoCmd(),
			saveCmd(),
			loadCmd(),
		},
		Before: initCom,
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

		buf.WriteString(fmt.Sprintf("%2s %2d: %s\n", current, tr.Index, tr.Title))
	}
	buf.WriteString(fmt.Sprintf("\n%d track(s)", len(ls)))
	fmt.Println(buf.String())
	return nil
}

func addCmd() *cli.Command {
	return &cli.Command{
		Name:        "add",
		Usage:       "add track(s) to the queue",
		ArgsUsage:   "TRACK...",
		Description: "TRACK can be a file or URL or - to read the list of tracks/URLs from stdin",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "append-only",
				Usage: "add but don't play",
			},
			&cli.BoolFlag{
				Name:  "replace",
				Aliases: []string{"r"},
				Usage: "add and play, stopping the current track",
			},
			&cli.BoolFlag{
				Name:  "next",
				Aliases: []string{"n"},
				Usage: "add as next track",
			},
		},
		Action: func(ctx *cli.Context) error {
			tracks := []string{ctx.Args().First()}
			tracks = append(tracks, ctx.Args().Tail()...)

			// check if we read from stdin
			if tracks[0] == "-" {
				tracks = []string{}
				in := bufio.NewScanner(os.Stdin)
				for in.Scan() {
					tracks = append(tracks, in.Text())
				}
				if err := in.Err(); err != nil {
					return fmt.Errorf("Error reading from stdin: %s", err)
				}
			}

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

func rmCmd() *cli.Command {
	return &cli.Command{
		Name:        "remove",
		Aliases:     []string{"rm"},
		Usage:       "remove tracks from the playlist",
		ArgsUsage:   "POSITION",
		Description: "POSITION can be a single index or a range expression like from..to (`to` is not removed)",
		Action: func(ctx *cli.Context) error {
			ns := strings.Split(ctx.Args().First(), "..")
			switch {
			// POSITION is range expression
			case len(ns) == 2:
				from, err := strconv.Atoi(ns[0])
				if err != nil {
					return err
				}
				to, err := strconv.Atoi(ns[1])
				if err != nil {
					return err
				}
				for i := 0; i < (to - from); i++ {
					if err := c.Remove(from); err != nil {
						return err
					}
				}
				return nil
			// POSITION is single index
			case len(ns) == 1:
				n, err := strconv.Atoi(ns[0])
				if err != nil {
					return err
				}
				return c.Remove(n)
			// POSITION is invalid
			default:
				return nil
			}
		},
	}
}

func clearCmd() *cli.Command {
	return &cli.Command{
		Name:  "clear",
		Usage: "remove all entries from the queue",
		Action: func(_ *cli.Context) error {
			return c.Clear()
		},
	}
}

func moveCmd() *cli.Command {
	return &cli.Command{
		Name:      "move",
		Aliases:   []string{"mv"},
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

func shuffleCmd() *cli.Command {
	return &cli.Command{
		Name:  "shuffle",
		Usage: "shuffle the current playlist",
		Action: func(_ *cli.Context) error {
			return c.Shuffle()
		},
	}
}

func nextCmd() *cli.Command {
	return &cli.Command{
		Name:  "next",
		Usage: "skips to the next track",
		Action: func(_ *cli.Context) error {
			return c.Next()
		},
	}
}

func prevCmd() *cli.Command {
	return &cli.Command{
		Name:  "prev",
		Usage: "skips to the previous track",
		Action: func(_ *cli.Context) error {
			return c.Prev()
		},
	}
}

func gotoCmd() *cli.Command {
	return &cli.Command{
		Name:      "goto",
		Usage:     "start playing NUMBER track on the queue",
		ArgsUsage: "NUMBER",
		Action: func(ctx *cli.Context) error {
			pos, err := strconv.Atoi(ctx.Args().First())
			if err != nil {
				return err
			}
			return c.Goto(pos)
		},
	}
}

func saveCmd() *cli.Command {
	return &cli.Command{
		Name:      "save",
		Usage:     "save the current playlist to a file",
		ArgsUsage: "PATH",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Aliases: []string{"f"},
				Usage: "playlist format to save as",
				Value: "m3u",
			},
		},
		Action: func(ctx *cli.Context) error {
			path := ctx.Args().First()
			if path == "" {
				return fmt.Errorf("Please specify a path.")
			}

			// tracks, err := c.List()
			// if err != nil {
			//	return err
			// }
			// return playlist.Save(path)
			return fmt.Errorf("not implemented yet")
		},
	}
}

func loadCmd() *cli.Command {
	return &cli.Command{
		Name:      "load",
		Usage:     "load playlist from a file",
		ArgsUsage: "PATH",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "append-only",
				Usage: "add but don't play",
			},
			&cli.BoolFlag{
				Name:  "replace",
				Aliases: []string{"r"},
				Usage: "add and play, stopping the current track",
			},
			&cli.BoolFlag{
				Name:  "next",
				Aliases: []string{"n"},
				Usage: "add as next track",
			},
		},
		Action: func(ctx *cli.Context) error {
			path := ctx.Args().First()
			if path == "" {
				return fmt.Errorf("Please specify a path.")
			}

			switch {
			case ctx.Bool("replace"):
				return c.LoadListReplace(path)
			case ctx.Bool("next"):
				return fmt.Errorf("not implemented yet")
				// return c.LoadNext(tracks...)
			default:
				return c.LoadListAppend(path)
			}
		},
	}
}
