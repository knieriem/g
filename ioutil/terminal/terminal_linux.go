// Copyright 2011 The Go Authors. All rights reserved.
// Copyright 2012 The g Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// IsTerminal is based on a function in go.crypto/ssh/terminal/util.go

package terminal

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

// These constants are declared here, rather than importing
// them from the syscall package as some syscall/unix packages, even
// on linux, for example gccgo, do not declare them.
const ioctlReadTermios = 0x5401  // unix.TCGETS
const ioctlWriteTermios = 0x5402 // unix.TCSETS

type FileDescriptor interface {
	Fd() uintptr
}

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(f FileDescriptor) bool {
	var termios unix.Termios
	_, _, err := unix.Syscall6(unix.SYS_IOCTL, f.Fd(), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
