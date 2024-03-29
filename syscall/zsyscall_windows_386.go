// Code generated by 'go generate'; DO NOT EDIT.

package syscall

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procRegEnumValueW       = modadvapi32.NewProc("RegEnumValueW")
	procCreateEventW        = modkernel32.NewProc("CreateEventW")
	procEscapeCommFunction  = modkernel32.NewProc("EscapeCommFunction")
	procFlushFileBuffers    = modkernel32.NewProc("FlushFileBuffers")
	procGetCommState        = modkernel32.NewProc("GetCommState")
	procGetOverlappedResult = modkernel32.NewProc("GetOverlappedResult")
	procSetCommState        = modkernel32.NewProc("SetCommState")
	procSetCommTimeouts     = modkernel32.NewProc("SetCommTimeouts")
	procSetConsoleMode      = modkernel32.NewProc("SetConsoleMode")
	procSetEvent            = modkernel32.NewProc("SetEvent")
	procSetupComm           = modkernel32.NewProc("SetupComm")
)

func RegEnumValue(h syscall.Handle, index uint32, vName *uint16, vNameLen *uint32, reserved *uint32, typ *uint32, data *byte, sz *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procRegEnumValueW.Addr(), 8, uintptr(h), uintptr(index), uintptr(unsafe.Pointer(vName)), uintptr(unsafe.Pointer(vNameLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(typ)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(sz)), 0)
	if r1 != ERROR_SUCCESS {
		err = errnoErr(e1)
	}
	return
}

func CreateEventW(sa *syscall.SecurityAttributes, manualReset int, initialState int, name *uint16) (hEv syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(sa)), uintptr(manualReset), uintptr(initialState), uintptr(unsafe.Pointer(name)), 0, 0)
	hEv = syscall.Handle(r0)
	if hEv == 0 {
		err = errnoErr(e1)
	}
	return
}

func EscapeCommFunction(h syscall.Handle, fn uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procEscapeCommFunction.Addr(), 2, uintptr(h), uintptr(fn), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func FlushFileBuffers(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFlushFileBuffers.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getCommState(h syscall.Handle, dcb *DCB) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommState.Addr(), 2, uintptr(h), uintptr(unsafe.Pointer(dcb)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetOverlappedResult(h syscall.Handle, ov *syscall.Overlapped, done *uint32, bWait int) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetOverlappedResult.Addr(), 4, uintptr(h), uintptr(unsafe.Pointer(ov)), uintptr(unsafe.Pointer(done)), uintptr(bWait), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setCommState(h syscall.Handle, dcb *DCB) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommState.Addr(), 2, uintptr(h), uintptr(unsafe.Pointer(dcb)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetCommTimeouts(h syscall.Handle, cto *CommTimeouts) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommTimeouts.Addr(), 2, uintptr(h), uintptr(unsafe.Pointer(cto)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetConsoleMode(h syscall.Handle, mode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetConsoleMode.Addr(), 2, uintptr(h), uintptr(mode), 0)
	if r1 == FALSE {
		err = errnoErr(e1)
	}
	return
}

func SetEvent(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetEvent.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetupComm(h syscall.Handle, inQSize uint32, outQSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupComm.Addr(), 3, uintptr(h), uintptr(inQSize), uintptr(outQSize))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
