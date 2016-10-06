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
	"unsafe"
)

const (
	termInFilename  = "/dev/tty"
	termOutFilename = "/dev/tty"
)

// These constants are declared here, rather than importing
// them from the syscall package as some syscall/unix packages, even
// on linux, for example gccgo, do not declare them.
const ioctlReadTermios = 0x5401  // unix.TCGETS
const ioctlWriteTermios = 0x5402 // unix.TCSETS

// State contains the state of a terminal.
type State struct {
	termios unix.Termios
}

type FileDescriptor interface {
	Fd() uintptr
}

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd FileDescriptor) bool {
	var termios unix.Termios
	_, _, err := unix.Syscall6(unix.SYS_IOCTL, fd.Fd(), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
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
func DisableFlags(fd FileDescriptor, flags Flag) (*State, error) {
	var oldState State
	if _, _, err := unix.Syscall6(unix.SYS_IOCTL, fd.Fd(), ioctlReadTermios, uintptr(unsafe.Pointer(&oldState.termios)), 0, 0, 0); err != 0 {
		return nil, err
	}

	newState := oldState.termios
	newState.Lflag &^= uint32(flags)
	if _, _, err := unix.Syscall6(unix.SYS_IOCTL, fd.Fd(), ioctlWriteTermios, uintptr(unsafe.Pointer(&newState)), 0, 0, 0); err != 0 {
		return nil, err
	}

	return &oldState, nil
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func Restore(fd FileDescriptor, state *State) error {
	_, _, err := unix.Syscall6(unix.SYS_IOCTL, fd.Fd(), ioctlWriteTermios, uintptr(unsafe.Pointer(&state.termios)), 0, 0, 0)
	return err
}
