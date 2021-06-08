// mksyscall.pl -tags linux,arm syscall_linux.go
// Code generated by the command above; DO NOT EDIT.

// +build linux,arm

package syscall

import "syscall"

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func ioctl(fd int, req uint, arg uintptr) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
