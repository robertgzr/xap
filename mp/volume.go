package mp

type Volume float64

func (c *Mp) Volume() (Volume, error) {
	vol, err := c.GetFloatProperty("volume")
	if err != nil {
		return Volume(0.0), err
	}
	return Volume(vol), nil
}

func (c *Mp) VolumeUp(val float64) error {
	vol, err := c.Volume()
	if err != nil {
		return err
	}
	newvol := float64(vol) + val
	if err := c.SetProperty("volume", newvol); err != nil {
		return err
	}
	return nil
}

func (c *Mp) VolumeDown(val float64) error {
	vol, err := c.Volume()
	if err != nil {
		return err
	}
	newvol := float64(vol) - val
	if err := c.SetProperty("volume", newvol); err != nil {
		return err
	}
	return nil
}
