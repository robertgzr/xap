package mp

import (
	"errors"
	"path"
	"path/filepath"
	"strconv"

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

func (c *Mp) Queue() (Queue, error) {
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

func (c *Mp) PlaylistPos() (int, error) {
	return c.GetIntProperty("playlist-pos")
}

func (c *Mp) PlaylistLen() (int, error) {
	return c.GetIntProperty("playlist-count")
}

func (c *Mp) LoadAppend(tracks ...string) error {
	return c.load(tracks, mpv.LoadFileModeAppend)
}

func (c *Mp) LoadPlay(tracks ...string) error {
	return c.load(tracks, mpv.LoadFileModeAppendPlay)
}

func (c *Mp) LoadReplace(tracks ...string) error {
	return c.load(tracks, mpv.LoadFileModeReplace)
}

func (c *Mp) LoadNext(tracks ...string) error {
	if err := c.load(tracks, mpv.LoadFileModeAppend); err != nil {
		return err
	}

	len, _ := c.PlaylistLen()
	pos, _ := c.PlaylistPos()
	return c.Move(len-1, pos+1)
}

func (c *Mp) load(tracks []string, mode string) error {
	for _, t := range tracks {
		// load file or URL
		if ignored(t) {
			continue
		}
		if err := c.loadSingleTrack(t, mode); err != nil {
			return errors.New(ErrLoadingTrack + err.Error())
		}
	}
	return nil
}

func ignored(filename string) bool {
	ignoredExt := map[string]struct{}{
		"nfo": struct{}{},
		"jpg": struct{}{},
		"png": struct{}{},
	}
	ext := filepath.Ext(filename)
	if _, ok := ignoredExt[ext]; ok {
		return true
	}
	return false
}

func (c *Mp) loadSingleTrack(track, mode string) error {
	if track == "" {
		return ErrNoFilepath
	}
	return c.Loadfile(track, mode)
}

func (c *Mp) LoadListAppend(path string) error {
	return c.loadlist(path, mpv.LoadListModeAppend)
}

func (c *Mp) LoadListReplace(path string) error {
	return c.loadlist(path, mpv.LoadListModeReplace)
}

// TODO: LoadListNext

func (c *Mp) loadlist(path, mode string) error {
	_, err := c.Exec("loadlist", path, mode)
	if err != nil {
		return err
	}
	return nil
}

func (c *Mp) Next() error {
	return c.PlaylistNext()
}

func (c *Mp) Prev() error {
	return c.PlaylistPrevious()
}

func (c *Mp) Move(from, to int) error {
	fromStr, toStr := strconv.Itoa(from), strconv.Itoa(to)
	_, err := c.Exec("playlist-move", fromStr, toStr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Mp) Shuffle() error {
	_, err := c.Exec("playlist-shuffle")
	if err != nil {
		return err
	}
	return nil
}

func (c *Mp) Remove(n int) error {
	_, err := c.Exec("playlist-remove", strconv.Itoa(n))
	if err != nil {
		return err
	}
	return nil
}

func (c *Mp) Clear() error {
	_, err := c.Exec("playlist-clear")
	if err != nil {
		return err
	}
	return nil
}

func (c *Mp) Goto(pos int) error {
	if len, err := c.PlaylistLen(); err != nil {
		return err
	} else if len < pos {
		return errors.New("Outside of playlist")
	}
	return c.SetProperty("playlist-pos", pos)
}
