// Some utility functions from SetupAPI.
// (see "Public Device Installation Functions",
//
//	http://msdn.microsoft.com/en-us/library/ff549791.aspx)
package setupapi

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// http://support.microsoft.com/kb/259695/en-us

var GuidSerialPorts = &Guid{
	Data1: 0x4D36E978,
	Data2: 0xE325,
	Data3: 0x11CE,
	Data4: [8]uint8{
		0xBF, 0xC1, 0x08, 0x00, 0x2B, 0xE1, 0x03, 0x18,
	},
}

//sys SetupDiGetClassDevs(class *Guid, enum *uint16, parent syscall.Handle, flags uint32) (devInfoSet syscall.Handle, err error) [failretval==syscall.InvalidHandle] = setupapi.SetupDiGetClassDevsW

//sys SetupDiGetDeviceRegistryProperty(devInfoSet syscall.Handle, diData *SpDevinfoData, prop uint32, regDataType *uint32, buf []byte, size *uint32) (err error) [failretval == 0] = setupapi.SetupDiGetDeviceRegistryPropertyW

//sys SetupDiEnumDeviceInfo(devInfoSet syscall.Handle, index uint32, diData *SpDevinfoData) (err error) [failretval == 0] = setupapi.SetupDiEnumDeviceInfo

//sys SetupDiCreateDeviceInfo(devInfoSet syscall.Handle, devName *uint16, g *Guid, devDesc *uint16, hwnd uintptr, cflags uint32, dataOut *SpDevinfoData) (err error) = setupapi.SetupDiCreateDeviceInfoW

//sys SetupDiCreateDeviceInfoList(g *Guid, hwnd uintptr) (devInfoSet syscall.Handle, err error) [failretval==syscall.InvalidHandle] = setupapi.SetupDiCreateDeviceInfoList

//sys SetupDiSetDeviceRegistryProperty(devInfoSet syscall.Handle, data *SpDevinfoData, prop uint32, buf *byte, sz uint32) (err error) = setupapi.SetupDiSetDeviceRegistryPropertyW

//sys SetupDiCallClassInstaller(installFn uintptr, devInfoSet syscall.Handle, data *SpDevinfoData) (err error) = setupapi.SetupDiCallClassInstaller

//sys SetupDiDestroyDeviceInfoList(devInfoSet syscall.Handle) (err error) = setupapi.SetupDiDestroyDeviceInfoList

//sys SetupDiGetINFClass(infPath *uint16, guid *Guid, className []uint16, reqSz *uint32) (err error) = setupapi.SetupDiGetINFClassW

//sys SetupDiOpenDevRegKey(devInfoSet syscall.Handle, diData *SpDevinfoData, scope uint32, hwProfile uint32, keyType uint32, desiredAccess uint32) (h syscall.Handle, err error) [failretval==syscall.InvalidHandle] = setupapi.SetupDiOpenDevRegKey

//sys SetupDiGetDeviceInstanceId(devInfoSet syscall.Handle, diData *SpDevinfoData, id []uint16, reqSz *uint32) (err error) = setupapi.SetupDiGetDeviceInstanceIdW

//sys SetupDiOpenDeviceInfo(devInfoSet syscall.Handle, deviceInstanceId *uint16, hwndParent syscall.Handle, openFlags uint32, diData *SpDevinfoData) (err error) [failretval == 0] = setupapi.SetupDiOpenDeviceInfoW

//sys SetupDiGetDeviceProperty(devInfoSet syscall.Handle, diData *SpDevinfoData, pKey *DevPropKey, pType *DevPropType, buf []byte, size *uint32, flags uint32) (err error) [failretval == 0] = setupapi.SetupDiGetDevicePropertyW

func NewDevinfoData() *SpDevinfoData {
	d := new(SpDevinfoData)
	d.CbSize = SpDevinfoDataSz
	return d
}

func NewDevPropKey(guid *Guid, pid uint32) *DevPropKey {
	return &DevPropKey{
		Fmtid: *guid,
		Pid:   pid,
	}
}

func NewGuidFromString(s string) (*Guid, error) {
	return NewGuidFromBytes([]byte(s))
}

func NewGuidFromBytes(b []byte) (*Guid, error) {
	b = bytes.Trim(b, "{}")

	if len(b) == 36 {
		b = bytes.ReplaceAll(b, []byte("-"), []byte(""))
	}
	if len(b) != 32 {
		return nil, fmt.Errorf("invalid guid format/length")
	}

	d := [8]byte{}
	if _, err := hex.Decode(d[0:4], b[0:8]); err != nil {
		return nil, err
	}
	data1 := uint32(d[3]) | uint32(d[2])<<8 | uint32(d[1])<<16 | uint32(d[0])<<24

	if _, err := hex.Decode(d[0:2], b[8:12]); err != nil {
		return nil, err
	}
	data2 := uint16(d[1]) | uint16(d[0])<<8

	if _, err := hex.Decode(d[0:2], b[12:16]); err != nil {
		return nil, err
	}
	data3 := uint16(d[1]) | uint16(d[0])<<8

	if _, err := hex.Decode(d[0:8], b[16:]); err != nil {
		return nil, err
	}

	return &Guid{
		Data1: data1,
		Data2: data2,
		Data3: data3,
		Data4: d,
	}, nil
}
