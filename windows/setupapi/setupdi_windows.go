package setupapi

import (
	"bytes"
	"encoding/binary"
	"syscall"

	"github.com/knieriem/g/windows/registry"
)

type DevinfoSet struct {
	h    syscall.Handle
	data *SpDevinfoData
	buf  []byte
}

func NewDevinfoSet(h syscall.Handle) *DevinfoSet {
	return &DevinfoSet{
		h:    h,
		data: NewDevinfoData(),
	}
}

func (di *DevinfoSet) Enum(i int) bool {
	return SetupDiEnumDeviceInfo(di.h, uint32(i), di.data) == nil
}

type Property []byte

func (p Property) String() string {
	data := []byte(p)
	if len(p) < 2 {
		return ""
	}
	ubuf := make([]uint16, len(data)/2)
	binary.Read(bytes.NewBuffer(data), byteOrder(isLittleEndian), ubuf)
	return syscall.UTF16ToString(ubuf)
}

func (p Property) Strings() (list []string) {
	data := []byte(p)
	if len(p) < 2 {
		return
	}
	ubuf := make([]uint16, len(data)/2)
	binary.Read(bytes.NewBuffer(data), byteOrder(isLittleEndian), ubuf)

	iPrev := 0
	for i := 0; i < len(ubuf); i++ {
		if ubuf[i] == 0 {
			if i-iPrev >= 2 {
				list = append(list, syscall.UTF16ToString(ubuf[iPrev:i]))
			}
			iPrev = i + 1
		}
	}
	return
}

func (di *DevinfoSet) DeviceRegistryProperty(prop uint32) Property {
	var bufSize uint32

	SetupDiGetDeviceRegistryProperty(di.h, di.data, prop, nil, nil, &bufSize)

	if cap(di.buf) < int(bufSize) {
		di.buf = make([]byte, bufSize, 2*bufSize)
	} else {
		di.buf = di.buf[:bufSize]
	}
	SetupDiGetDeviceRegistryProperty(di.h, di.data, prop, nil, di.buf, nil)
	return Property(di.buf)
}

func (di *DevinfoSet) OpenDevRegKey() (k *registry.Key, err error) {
	h, err := SetupDiOpenDevRegKey(di.h, di.data, DICS_FLAG_GLOBAL, 0, DIREG_DEV, syscall.KEY_READ)
	if err == nil {
		k = &registry.Key{h}
	}
	return
}

func (di *DevinfoSet) DeviceInstanceID() (id string, err error) {
	var bufSize uint32

	SetupDiGetDeviceInstanceId(di.h, di.data, nil, &bufSize)
	ubuf := make([]uint16, bufSize)
	err = SetupDiGetDeviceInstanceId(di.h, di.data, ubuf, nil)
	if err == nil {
		id = syscall.UTF16ToString(ubuf)
	}
	return
}

const (
	isLittleEndian = syscall.REG_DWORD == syscall.REG_DWORD_LITTLE_ENDIAN
)

func byteOrder(little bool) binary.ByteOrder {
	if little {
		return binary.LittleEndian
	}
	return binary.BigEndian
}
