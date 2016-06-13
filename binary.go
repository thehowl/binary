// Package binary approaches binary encoding with a different approach from
// encoding/binary. The main aim of this package is to be *fast*, and thus not
// to use reflection, and with more type-safety.
//
// By its nature, this package is statically typed, and as such will not
// support arrays, structs and slices of non-builtin types (or slices of
// slices). Apart from that, behaviour is pretty much like that of
// encoding/binary.
package binary

import (
	encodingBinary "encoding/binary"
	"errors"
	"io"
	"strconv"
)

var errChainNil = errors.New("thehowl/binary: Chain is nil")

// Brought from the encoding/binary package as a shorthand.
var (
	LittleEndian = encodingBinary.LittleEndian
	BigEndian    = encodingBinary.BigEndian
)

// WriteChain wraps around an io.Writer and a encoding.ByteOrder, and writes data
// into the Writer using the given ByteOrder. WriteChain is not thread-safe.
type WriteChain struct {
	Writer    io.Writer
	ByteOrder encodingBinary.ByteOrder
	written   int
	err       error
}

// End finishes writing to the Writer, and returns the amount of written bytes
// and any eventual error occured during the WriteChain lifetime. It also clears out
// the written bytes and the error, making WriteChain usable as defined by the user
// initially again.
func (c *WriteChain) End() (int, error) {
	if c == nil {
		return -1, errChainNil
	}
	written := c.written
	c.written = 0
	err := c.err
	c.err = nil
	return written, err
}

func checkWritten(written, expected int) error {
	if written != expected {
		return errors.New("Expected to write " + strconv.Itoa(expected) + " bytes, wrote " + strconv.Itoa(written))
	}
	return nil
}

// Reader is the simplified version of ReadChain. Instead of having to pass
// pointers to decode bytes, using Reader you can read data directly and
// do, for instance, simple assignments.
type Reader struct {
	Reader    io.Reader
	ByteOrder encodingBinary.ByteOrder
	read      int
	err       error
}

// End returns the amount of read bytes and any eventual error, and sets the
// internal amount of written bytes to 0, and the error to nil.
func (r *Reader) End() (int, error) {
	if r == nil {
		return -1, errChainNil
	}
	read := r.read
	r.read = 0
	err := r.err
	r.err = nil
	return read, err
}
