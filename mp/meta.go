package mp

import (
	"fmt"
	"strings"
	"time"

	"github.com/blang/mpv"
)

type Metadata struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Track  string `json:"track"`
	Date   string `json:"date"`

	Pos Position `json:"position"`
}

type Position struct {
	Len         time.Duration `json:"duration"`
	Current     time.Duration `json:"current"`
	CurrentPerc float64       `json:"current_percent"`
}

func (c *Mp) Now() (meta Metadata, err error) {
	var res = new(mpv.Response)
	res, err = c.Exec("get_property", "metadata")
	if err != nil {
		return meta, err
	}

	data, ok := res.Data.(map[string]interface{})
	if !ok {
		return meta, mpv.ErrInvalidType
	}
	if _, ok := data["title"]; ok {
		meta.Title, ok = data["title"].(string)
	}
	if _, ok := data["artist"]; ok {
		meta.Artist = data["artist"].(string)
	}
	if _, ok := data["album"]; ok {
		meta.Album = data["album"].(string)
	}
	if _, ok := data["genre"]; ok {
		meta.Genre = data["genre"].(string)
	}
	if _, ok := data["track"]; ok {
		meta.Track = data["track"].(string)
	}
	if _, ok := data["date"]; ok {
		meta.Date = data["date"].(string)
	}

	if meta.Title == "" {
		title, err := c.MediaTitle()
		if err != nil {
			return meta, err
		}
		meta.Title = title
	}
	meta.Title = strings.TrimSuffix(meta.Title, "\"")
	meta.Title = strings.TrimPrefix(meta.Title, "\"")

	var dur float64
	dur, err = c.Duration()
	if err != nil {
		return
	}
	meta.Pos.Len, err = time.ParseDuration(fmt.Sprintf("%fs", dur))
	if err != nil {
		return
	}
	pos, err := c.Position()
	if err != nil {
		return
	}
	meta.Pos.Current, err = time.ParseDuration(fmt.Sprintf("%fs", pos))
	if err != nil {
		return
	}
	meta.Pos.CurrentPerc, err = c.PercentPosition()
	if err != nil {
		return
	}
	return
}

func (c *Mp) MediaTitle() (string, error) {
	return c.GetProperty("media-title")
}
