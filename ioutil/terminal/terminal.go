package terminal

import (
	"os"
)

// OpenInput tries to get read and write access to the process' terminal.
func OpenInput() (*os.File, error) {
	return os.OpenFile(termInFilename, os.O_RDONLY, 0)
}

// OpenOutput tries to get write access to the process' terminal.
func OpenOutput() (*os.File, error) {
	return os.OpenFile(termOutFilename, os.O_WRONLY, 0)
}
