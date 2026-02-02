//go:build windows

// Package tmux provides a wrapper for tmux session operations via subprocess.
// Note: tmux does not exist on Windows, so this file provides no-op stubs
// to allow cross-compilation. The actual tmux operations will fail at runtime
// on Windows, but the code will compile.
package tmux

// Signal represents a process signal (stub for Windows compatibility).
type Signal int

// signalSIGTERM is a stub for SIGTERM on Windows.
var signalSIGTERM Signal = 15

// signalSIGKILL is a stub for SIGKILL on Windows.
var signalSIGKILL Signal = 9

// killProcessGroup is a no-op on Windows since tmux doesn't exist there.
// This allows the code to compile for Windows cross-compilation testing.
func killProcessGroup(pgid int, sig Signal) error {
	// No-op: tmux and Unix process groups don't exist on Windows.
	// Any code path that reaches here on Windows would have already failed
	// when trying to invoke tmux commands.
	return nil
}
