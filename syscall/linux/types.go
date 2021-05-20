// +build ignore

package syscall

/*
#include <linux/serial.h>
*/
import "C"

type Serial C.struct_serial_struct
