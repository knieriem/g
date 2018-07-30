// +build ignore

package syscall

/*
#include <termios.h>
#include <unistd.h>
#include <linux/serial.h>
*/
import "C"

type Termios	C.struct_termios
type Int	C.int

type Serial	C.struct_serial_struct
