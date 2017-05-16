package com

import (
	"github.com/blang/mpv"
)

type Metadata struct {
	Title  string
	Artist string
	Album  string
	Genre  string
	Nr     string
	Date   string

	Duration float64
}

func (c *Com) CurrentTrack() (Metadata, error) {
	var meta Metadata
	res, err := c.Exec("get_property", "metadata")
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
		meta.Nr = data["track"].(string)
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

	d, err := c.Duration()
	if err != nil {
		return meta, err
	}
	meta.Duration = d

	return meta, err
}

func (c *Com) MediaTitle() (string, error) {
	return c.GetProperty("media-title")
}
