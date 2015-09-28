package terminal

import (
	"errors"
	"io"
	"os"
)

type Term struct {
	io.Reader
	io.WriteCloser
}

// Open tries to get read and write access to the process'
// terminal, in case os.Stdin is linked with the terminal.
func Open() (t *Term, err error) {
	if !IsTerminal(os.Stdin) {
		err = errors.New("os.Stdin is not connected to a terminal")
		return
	}
	f, err := os.OpenFile(termOutFilename, os.O_WRONLY, 0)
	if err != nil {
		return
	}
	t = new(Term)
	t.Reader = os.Stdin
	t.WriteCloser = f
	return
}
