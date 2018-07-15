// Package binary approaches binary encoding with a different approach from
// encoding/binary. The main aim of this package is to be *fast*, and thus not
// to use reflection, and with more type-safety.
//
// By its nature, this package is statically typed, and as such will not
// support arrays, structs and slices of compound/non-builtin types (or slices
// of slices). Apart from that, behaviour is pretty much like that of
// encoding/binary.
package binary

import (
	encodingBinary "encoding/binary"
	"errors"
	"io"
)

var errChainNil = errors.New("thehowl/binary: WriteChain is nil")

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
	_buf      []byte
}

// Write is a simple wrapper around c.Writer.Write. The WriteChain's internal
// written counter will be incremented and the error will be set if any, so that
// c.End returns the correct result.
func (c *WriteChain) Write(p []byte) (int, error) {
	if c == nil {
		return 0, errChainNil
	}
	if c.err != nil {
		return 0, c.err
	}
	i, err := c.Writer.Write(p)
	c.written += i
	if err != nil {
		c.err = err
	}
	return i, err
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
	if c._buf != nil {
		bufpool.Put(c._buf)
		c._buf = nil
	}
	return written, err
}

// Reader is the simplified version of ReadChain. Instead of having to pass
// pointers to decode bytes, using Reader you can read data directly and
// do, for instance, simple assignments.
type Reader struct {
	Reader    io.Reader
	ByteOrder encodingBinary.ByteOrder
	read      int
	err       error
	buf       []byte
}

var errReaderNil = errors.New("thehowl/binary: Reader is nil")

// Read is a simple wrapper around r.Reader.Read. The Reader's internal
// read counter will be incremented and the error will be set if any, so that
// r.End returns the correct result.
func (r *Reader) Read(p []byte) (int, error) {
	if r == nil {
		return 0, errReaderNil
	}
	if r.err != nil {
		return 0, r.err
	}
	i, err := r.Reader.Read(p)
	r.read += i
	if err != nil {
		r.err = err
	}
	return i, err
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
