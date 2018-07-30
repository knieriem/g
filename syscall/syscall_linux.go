package syscall

import (
	"syscall"
)

//sys	Fcntl(fd uintptr, cmd int, arg int) (val int, err error)

//sys	IoctlTermios(fd uintptr, action int, t *Termios) (err error) = SYS_IOCTL
//sys	IoctlModem(fd uintptr, action int, flags *Int) (err error) = SYS_IOCTL
//sys	IoctlSerial(fd uintptr, action int, s *Serial) (err error) = SYS_IOCTL

func (t *Termios) SetInSpeed(s int) {
	//	t.Iflag = t.Iflag&^CBAUD | uint32(s)&CBAUD
}
func (t *Termios) SetOutSpeed(s int) {
	t.Cflag = t.Cflag&^CBAUD | uint32(s)&CBAUD
}

// Copied from syscall/syscall_unix.go

// Do the interface allocations only once for common
// Errno values.
var (
	errEAGAIN error = syscall.EAGAIN
	errEINVAL error = syscall.EINVAL
	errENOENT error = syscall.ENOENT
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case syscall.EAGAIN:
		return errEAGAIN
	case syscall.EINVAL:
		return errEINVAL
	case syscall.ENOENT:
		return errENOENT
	}
	return e
}
