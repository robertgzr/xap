package com

import (
	"strings"

	"github.com/blang/mpv"
)

type AudioDevice struct {
	ID          int
	Name        string
	Description string
	Current     bool
}

func (c *Com) AudioDeviceList() ([]AudioDevice, error) {
	var adlst []AudioDevice

	current, err := c.GetAudioDevice()
	if err != nil {
		return adlst, err
	}

	res, err := c.Exec("get_property", "audio-device-list")
	if err != nil {
		return nil, err
	}

	rawLst, ok := res.Data.([]interface{})
	if !ok {
		return nil, mpv.ErrInvalidType
	}

	if len(rawLst) == 0 {
		return adlst, nil
	}

	for i, el := range rawLst {
		entry := el.(map[string]interface{})
		device := AudioDevice{
			ID:          i,
			Name:        entry["name"].(string),
			Description: entry["description"].(string),
			Current:     false,
		}
		if device.Name == current {
			device.Current = true
		}
		adlst = append(adlst, device)
	}

	return adlst, err
}

func (c *Com) SetAudioDevice(id int) error {
	adlst, err := c.AudioDeviceList()
	if err != nil {
		return err
	}
	return c.SetProperty("audio-device", adlst[id].Name)
}

func (c *Com) GetAudioDevice() (string, error) {
	ad, err := c.GetProperty("audio-device")
	if err != nil {
		return "", err
	}
	return strings.Trim(ad, "\""), nil
}