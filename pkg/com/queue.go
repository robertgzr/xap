package com

import (
	"errors"
	"path"
	"strconv"
	"strings"

	"github.com/blang/mpv"
)

var (
	ErrLoadingTrack string = "Error loading track: "
)

type Queue []Track

type Track struct {
	Index    int
	Title    string
	Location string
	Playing  bool
	Current  bool
}

func (c *Com) Queue() (Queue, error) {
	var plst Queue

	res, err := c.Exec("get_property", "playlist")
	if err != nil {
		return nil, err
	}

	rawLst, ok := res.Data.([]interface{})
	if !ok {
		return nil, mpv.ErrInvalidType
	}

	if len(rawLst) == 0 {
		return plst, nil
	}

	for i, el := range rawLst {
		var track Track = Track{Index: i}

		entry := el.(map[string]interface{})
		if cr, ok := entry["current"]; ok {
			track.Current = cr.(bool)
		}
		if pl, ok := entry["playing"]; ok {
			track.Playing = pl.(bool)
		}
		if fp, ok := entry["filename"]; ok {
			track.Location = fp.(string)
		}
		if tl, ok := entry["title"]; ok {
			track.Title = tl.(string)
		} else {
			track.Title = path.Base(track.Location)
		}

		plst = append(plst, track)
	}

	return plst, nil
}

func (c *Com) PlaylistPos() (int, error) {
	return c.GetIntProperty("playlist-pos")
}

func (c *Com) PlaylistLen() (int, error) {
	return c.GetIntProperty("playlist-count")
}

func (c *Com) LoadAppend(tracks ...string) error {
	return c.load(tracks, mpv.LoadFileModeAppend)
}

func (c *Com) LoadPlay(tracks ...string) error {
	return c.load(tracks, mpv.LoadFileModeAppendPlay)
}

func (c *Com) LoadReplace(tracks ...string) error {
	return c.load(tracks, mpv.LoadFileModeReplace)
}

func (c *Com) LoadNext(tracks ...string) error {
	if err := c.load(tracks, mpv.LoadFileModeAppend); err != nil {
		return err
	}

	len, _ := c.PlaylistLen()
	pos, _ := c.PlaylistPos()
	return c.Move(len-1, pos+1)
}

func (c *Com) load(tracks []string, mode string) error {
	for _, t := range tracks {
		// load file or URL
		switch {
		case strings.HasSuffix(t, ".nfo"):
			fallthrough
		case strings.HasSuffix(t, ".jpg"):
			fallthrough
		case strings.HasSuffix(t, ".png"):
			continue
		}
		if err := c.loadSingleTrack(t, mode); err != nil {
			return errors.New(ErrLoadingTrack + err.Error())
		}
	}
	return nil
}

func (c *Com) loadSingleTrack(track, mode string) error {
	if track == "" {
		return ErrNoFilepath
	}
	return c.Loadfile(track, mode)
}

func (c *Com) LoadListAppend(path string) error {
	return c.loadlist(path, mpv.LoadListModeAppend)
}

func (c *Com) LoadListReplace(path string) error {
	return c.loadlist(path, mpv.LoadListModeReplace)
}

// TODO: LoadListNext

func (c *Com) loadlist(path, mode string) error {
	_, err := c.Exec("loadlist", path, mode)
	if err != nil {
		return err
	}
	return nil
}

func (c *Com) Next() error {
	return c.PlaylistNext()
}

func (c *Com) Prev() error {
	return c.PlaylistPrevious()
}

func (c *Com) Move(from, to int) error {
	fromStr, toStr := strconv.Itoa(from), strconv.Itoa(to)
	_, err := c.Exec("playlist-move", fromStr, toStr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Com) Shuffle() error {
	_, err := c.Exec("playlist-shuffle")
	if err != nil {
		return err
	}
	return nil
}

func (c *Com) Remove(n int) error {
	_, err := c.Exec("playlist-remove", strconv.Itoa(n))
	if err != nil {
		return err
	}
	return nil
}

func (c *Com) Clear() error {
	_, err := c.Exec("playlist-clear")
	if err != nil {
		return err
	}
	return nil
}

func (c *Com) Goto(pos int) error {
	if len, err := c.PlaylistLen(); err != nil {
		return err
	} else if len < pos {
		return errors.New("Outside of playlist")
	}
	return c.SetProperty("playlist-pos", pos)
}
