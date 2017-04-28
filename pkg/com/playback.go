package com

import "strconv"

func (c *Com) Play() error {
	posString, err := c.GetProperty("playlist-pos")
	if err != nil {
		return err
	}
	if posString == "" {
		println("nopos")
		return ErrNoPlayerRunning
	}

	pos, err := strconv.Atoi(posString)
	if err != nil {
		return err
	}
	println(pos, pos+1)

	return c.SetPause(false)
}

func (c *Com) Stop() error {
	_, err := c.Exec("stop")
	return err

}
