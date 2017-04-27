package com

import (
	"errors"
	"os"

	"github.com/blang/mpv"
)

var (
	ErrNoPlayerRunning = errors.New("No player running")
	ErrNoFilepath      = errors.New("need to provide a filepath or URL")
)

// Com wraps mpv.Client
type Com struct {
	mpv.Client
}

// NewCom returns a new client on the socket at socketPath
func NewCom(socketPath string) (*Com, error) {
	if _, err := os.Stat(socketPath); err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoPlayerRunning
		}
		return nil, err
	}

	ipcc := mpv.NewIPCClient(socketPath)
	mpvc := mpv.NewClient(ipcc)

	return &Com{*mpvc}, nil
}

func (c *Com) Quit() error {
	_, err := c.Exec("quit")
	return err
}
