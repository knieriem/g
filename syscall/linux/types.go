// +build ignore

package syscall

/*
#include <linux/serial.h>
*/
import "C"

type Serial C.struct_serial_struct

type	SerialRS485 C.struct_serial_rs485
