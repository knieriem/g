// Copyright 2011 The Go Authors. All rights reserved.
// Copyright 2012 The g Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package terminal

import (
	"syscall"

	win "github.com/knieriem/g/syscall"
)

const (
	termInFilename  = "CONIN$"
	termOutFilename = "CONOUT$"
)

// IsTerminal and Restore are based on functions in
// golang.org/x/crypto/ssh/terminal/util_windows.go,
// and DisableFlags is derived from MakeRaw().

type FileDescriptor interface {
	Fd() uintptr
}

type State struct {
	mode uint32
}

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd FileDescriptor) (is bool) {
	var mode uint32
	is = syscall.GetConsoleMode(handle(fd), &mode) == nil
	return
}

const (
	enableLineInput      = 2
	enableEchoInput      = 4
	enableProcessedInput = 1
)

type Flag uint32

const (
	LineFlag   Flag = enableLineInput
	EchoFlag   Flag = enableEchoInput
	SignalFlag Flag = enableProcessedInput
)

// DisableFlags disables the terminal flags specified in the function
// argument 'flags', and returns the previous state of the terminal
// so that it can be restored.
func DisableFlags(fd FileDescriptor, flags Flag) (*State, error) {
	var st uint32
	err := syscall.GetConsoleMode(handle(fd), &st)
	if err != nil {
		return nil, err
	}
	st &^= uint32(flags)
	err = win.SetConsoleMode(handle(fd), st)
	return &State{st}, err
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func Restore(fd FileDescriptor, state *State) error {
	return win.SetConsoleMode(handle(fd), state.mode)
}

func handle(f FileDescriptor) syscall.Handle {
	return syscall.Handle(f.Fd())
}
