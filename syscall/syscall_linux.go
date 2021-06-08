package syscall

import (
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// IoctlSetSerial performs an ioctl on fd with a *Serial.
func IoctlSetSerial(fd int, value *Serial) error {
	err := ioctl(fd, unix.TIOCSSERIAL, uintptr(unsafe.Pointer(value)))
	runtime.KeepAlive(value)
	return err
}

func IoctlGetSerial(fd int) (*Serial, error) {
	var value Serial
	err := ioctl(fd, unix.TIOCGSERIAL, uintptr(unsafe.Pointer(&value)))
	return &value, err
}

// IoctlSetSerialRS485 performs an ioctl on fd with a *SerialRS485.
func IoctlSetSerialRS485(fd int, value *SerialRS485) error {
	err := ioctl(fd, unix.TIOCSRS485, uintptr(unsafe.Pointer(value)))
	runtime.KeepAlive(value)
	return err
}

func IoctlGetSerialRS485(fd int) (*SerialRS485, error) {
	var value SerialRS485
	err := ioctl(fd, unix.TIOCGRS485, uintptr(unsafe.Pointer(&value)))
	return &value, err
}

//sys	ioctl(fd int, req uint, arg uintptr) (err error)

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
