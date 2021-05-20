// Copyright 2011 The Go Authors. All rights reserved.
// Copyright 2012 The g Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// IsTerminal and Restore are based on functions in
// golang.org/x/crypto/ssh/terminal/util.go,
// and DisableFlags is derived from MakeRaw().

package terminal

import (
	"golang.org/x/sys/unix"
)

const (
	termInFilename  = "/dev/tty"
	termOutFilename = "/dev/tty"
)

// State contains the state of a terminal.
type State struct {
	termios unix.Termios
}

type FileDescriptor interface {
	Fd() uintptr
}

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(f FileDescriptor) bool {
	_, err := unix.IoctlGetTermios(int(f.Fd()), unix.TCGETS)
	return err == nil
}

type Flag uint32

const (
	LineFlag   Flag = unix.ICANON
	EchoFlag   Flag = unix.ECHO
	SignalFlag Flag = unix.ISIG
)

// DisableFlags disables the terminal flags specified in the function
// argument 'flags', and returns the previous state of the terminal
// so that it can be restored.
func DisableFlags(f FileDescriptor, flags Flag) (*State, error) {
	var old State

	fd := int(f.Fd())
	t, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return nil, err
	}
	old.termios = *t

	t.Lflag &^= uint32(flags)
	err = unix.IoctlSetTermios(fd, unix.TCSETS, t)
	if err != nil {
		return nil, err
	}

	return &old, nil
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func Restore(f FileDescriptor, state *State) error {
	return unix.IoctlSetTermios(int(f.Fd()), unix.TCSETS, &state.termios)
}
