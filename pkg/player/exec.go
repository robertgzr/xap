package player

import "os/exec"

func start() error {
	pl.proc = exec.Command("mpv", "--idle", "--input-ipc-server="+mpvSocket)
	return pl.proc.Start()
}
