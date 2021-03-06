package binary

import (
	"errors"
	"strconv"
)

// Uint8 decodes an uint8 from the Reader.
func (r *Reader) Uint8() uint8 {
	b := r.rd(1)
	if b == nil {
		return 0
	}
	return b[0]
}

// Uint16 decodes an uint16 from the Reader.
func (r *Reader) Uint16() uint16 {
	b := r.rd(2)
	if b == nil {
		return 0
	}
	return r.ByteOrder.Uint16(b)
}

// Uint32 decodes an uint32 from the Reader.
func (r *Reader) Uint32() uint32 {
	b := r.rd(4)
	if b == nil {
		return 0
	}
	return r.ByteOrder.Uint32(b)
}

// Uint64 decodes an uint64 from the Reader.
func (r *Reader) Uint64() uint64 {
	b := r.rd(8)
	if b == nil {
		return 0
	}
	return r.ByteOrder.Uint64(b)
}

// Byte decodes an uint64 from the Reader.
func (r *Reader) Byte() byte {
	return byte(r.Uint8())
}

// Bool decodes a bool from the Reader.
func (r *Reader) Bool() bool {
	return r.Uint8() > 0
}

// not read (present) because read (past participle) is in the struct's
// fields.
func (r *Reader) rd(amt int) []byte {
	if r == nil || r.err != nil {
		return nil
	}
	ret := make([]byte, amt)
	n, err := r.Reader.Read(ret)
	if err != nil {
		r.err = err
		return nil
	}
	if n != amt {
		r.err = errors.New("thehowl/binary: expected to read " + strconv.Itoa(amt) + " bytes, read " + strconv.Itoa(n))
		return nil
	}
	r.read += n
	return ret
}
