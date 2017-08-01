package com

func (c Com) Paused() bool {
	ok, err := c.Pause()
	if err != nil {
		return false
	}
	return ok
}

func (c *Com) Play() error {
	posString, err := c.GetProperty("playlist-pos")
	if err != nil {
		return err
	}
	if posString == "<nil>" {
		return c.SetProperty("playlist-pos", 0)
	}

	return c.SetPause(false)
}

func (c *Com) Stop() error {
	_, err := c.Exec("stop")
	return err

}
