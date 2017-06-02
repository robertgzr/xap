package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/blang/mpv"
	"github.com/urfave/cli"
)

func ControlCommands() []*cli.Command {
	return []*cli.Command{
		nowCmd(),
		playCmd(),
		pauseCmd(),
		stopCmd(),
		from0Cmd(),
	}
}

func durationFmt(d time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d", int(d.Hours())%24, int(d.Minutes())%60, int(d.Seconds())%60)
}

func nowCmd() *cli.Command {
	return &cli.Command{
		Name:     "now",
		Category: "control",
		Usage:    "show currently playing song",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			m, err := c.CurrentTrack()
			if err != nil {
				return err
			}

			dur, err := c.Duration()
			if err != nil {
				return err
			}
			tdur, err := time.ParseDuration(fmt.Sprintf("%fs", dur))
			if err != nil {
				return err
			}

			pos, err := c.Position()
			if err != nil {
				return err
			}
			tpos, err := time.ParseDuration(fmt.Sprintf("%fs", pos))
			if err != nil {
				return err
			}

			ppos, err := c.PercentPosition()
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			buf.WriteString("CURRENT:\n")
			buf.WriteString(fmt.Sprintf("| %v\n", m.Title))
			if m.Artist != "" {
				buf.WriteString(fmt.Sprintf("| %v\n", m.Artist))
			}
			if m.Album != "" {
				buf.WriteString(fmt.Sprintf("| %v (%v) %v\n", m.Album, m.Date, m.Nr))
			}

			buf.WriteString(fmt.Sprintf("\n%s / %s (%.1f%%)", durationFmt(tpos), durationFmt(tdur), ppos))
			fmt.Println(buf.String())
			return nil
		},
	}
}

func playCmd() *cli.Command {
	return &cli.Command{
		Name:     "play",
		Category: "control",
		Usage:    "start playing the current file",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.Play()
		},
	}
}

func from0Cmd() *cli.Command {
	return &cli.Command{
		Name:     "from0",
		Category: "control",
		Usage:    "restart playback of the current track",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.Seek(0, mpv.SeekModeAbsolute)
		},
	}
}

func pauseCmd() *cli.Command {
	return &cli.Command{
		Name:     "pause",
		Category: "control",
		Usage:    "pause the current file",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.SetPause(true)
		},
	}
}

func stopCmd() *cli.Command {
	return &cli.Command{
		Name:     "stop",
		Category: "control",
		Usage:    "stop the current file",
		Before:   initCom,
		Action: func(_ *cli.Context) error {
			return c.Stop()
		},
	}
}
