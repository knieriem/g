// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs linux/types.go

package syscall

type Termios struct {
	Iflag     uint32
	Oflag     uint32
	Cflag     uint32
	Lflag     uint32
	Line      uint8
	Cc        [32]uint8
	Pad_cgo_0 [3]byte
	Ispeed    uint32
	Ospeed    uint32
}
type Int int32

type Serial struct {
	Type            int32
	Line            int32
	Port            uint32
	Irq             int32
	Flags           int32
	Xmit_fifo_size  int32
	Custom_divisor  int32
	Baud_base       int32
	Close_delay     uint16
	Io_type         int8
	Reserved_char   [1]int8
	Hub6            int32
	Closing_wait    uint16
	Closing_wait2   uint16
	Pad_cgo_0       [4]byte
	Iomem_base      *uint8
	Iomem_reg_shift uint16
	Pad_cgo_1       [2]byte
	Port_high       uint32
	Iomap_base      uint64
}
