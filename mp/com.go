package mp

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/blang/mpv"
)

var (
	ErrNoPlayerRunning = errors.New("No player running")
	ErrNoFilepath      = errors.New("need to provide a filepath or URL")
)

// Mp wraps mpv.Client
type Mp struct {
	mpv.Client
}

// NewMp returns a new client on the socket at socketPath
func NewMp(socketPath string) (c *Mp, err error) {
	socketPath, err = filepath.EvalSymlinks(socketPath)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(socketPath); err != nil {
		return nil, err
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(ErrNoPlayerRunning)
			os.Exit(1)
		}
	}()

	if _, err := os.Stat(socketPath); err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoPlayerRunning
		}
		return nil, err
	}

	ipcc := mpv.NewIPCClient(socketPath)
	mpvc := mpv.NewClient(ipcc)

	// sanity check, do we have a connection...?
	res, err := ipcc.Exec("get_version")
	if res == nil {
		return nil, err
	}

	return &Mp{*mpvc}, nil
}

func (c *Mp) GetIntProperty(prop string) (int, error) {
	res, err := c.GetProperty(prop)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(res)
}

func (c *Mp) Quit() error {
	_, err := c.Exec("quit")
	return err
}
