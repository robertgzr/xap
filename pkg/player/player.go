package player

import (
	"os/exec"
	"syscall"

	"github.com/blang/mpv"
)

const mpvSocket string = "/tmp/mpv.sock"

// Player holds both the mpv process and the connected client
type Player struct {
	proc *exec.Cmd
	m    *mpv.Client
}

var pl *Player // the global player instance - we only want one

// NewPlayer starts the mpv process and initializes the client
// It returns the Player instance
func NewPlayer() (*Player, error) {
	// return instance if initialized
	if pl != nil {
		return pl, nil
	}

	if err := pl.init(); err != nil {
		return nil, err
	}

	return pl, nil
}

// Quit closes the client and terminates the mpv process
func (p *Player) Quit() error {
	p.m = nil
	return p.proc.Process.Signal(syscall.SIGTERM)
}

func (p *Player) init() error {
	if p.proc.ProcessState.Exited() {
		err := start()
		if err != nil {
			return err
		}
	}

	ipcc := mpv.NewIPCClient(mpvSocket)
	p.m = mpv.NewClient(ipcc)

	return p.defaultProperties()
}
