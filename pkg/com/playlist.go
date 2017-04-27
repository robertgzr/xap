package com

import (
	"bytes"
	"fmt"
)

func (c *Com) List() (string, error) {
	resp, err := c.Exec("get_property", "playlist")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	for i, el := range resp.Data.([]interface{}) {
		entry := el.(map[string]interface{})

		var current string
		if entry["current"] != nil && entry["current"].(bool) {
			current = ">"
		}
		buf.WriteString(fmt.Sprintf("%2s %d: %s\n", current, i, entry["filename"]))
	}

	return buf.String(), nil
}

func (c *Com) Load(fpath, mode string) error {
	if fpath == "" {
		return ErrNoFilepath
	}
	// TODO: check for allowd modes [append, append-play, replace]
	return c.Loadfile(fpath, mode)
}

func (c *Com) Next() error {
	return c.PlaylistNext()
}

func (c *Com) Prev() error {
	return c.PlaylistPrevious()
}
