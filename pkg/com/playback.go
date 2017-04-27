package com

func (c *Com) Stop() error {
	_, err := c.Exec("stop")
	return err

}
