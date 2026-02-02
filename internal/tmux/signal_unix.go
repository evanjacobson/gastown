//go:build !windows

// Package tmux provides a wrapper for tmux session operations via subprocess.
package tmux

import "syscall"

// killProcessGroup sends a signal to a process group.
// Uses syscall.Kill with a negative PID to target the process group (POSIX).
func killProcessGroup(pgid int, sig syscall.Signal) error {
	return syscall.Kill(-pgid, sig)
}

// signalSIGTERM is the SIGTERM signal for graceful termination.
var signalSIGTERM = syscall.SIGTERM

// signalSIGKILL is the SIGKILL signal for forceful termination.
var signalSIGKILL = syscall.SIGKILL
