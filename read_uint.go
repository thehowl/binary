package binary

import (
	"errors"
	"strconv"
)

// Uint8 decodes an uint8 from the Reader.
func (r *ReadChain) Uint8(u *uint8) *ReadChain {
	*u = r.uint8()
	return r
}

func (r *ReadChain) uint8() uint8 {
	b := r.rd(1)
	if b == nil {
		return 0
	}
	return b[0]
}

// Uint16 decodes an uint16 from the Reader.
func (r *ReadChain) Uint16(u *uint16) *ReadChain {
	*u = r.uint16()
	return r
}

func (r *ReadChain) uint16() uint16 {
	b := r.rd(2)
	if b == nil {
		return 0
	}
	return r.ByteOrder.Uint16(b)
}

// Uint32 decodes an uint32 from the Reader.
func (r *ReadChain) Uint32(u *uint32) *ReadChain {
	*u = r.uint32()
	return r
}

func (r *ReadChain) uint32() uint32 {
	b := r.rd(4)
	if b == nil {
		return 0
	}
	return r.ByteOrder.Uint32(b)
}

// Uint64 decodes an uint64 from the Reader.
func (r *ReadChain) Uint64(u *uint64) *ReadChain {
	*u = r.uint64()
	return r
}

func (r *ReadChain) uint64() uint64 {
	b := r.rd(8)
	if b == nil {
		return 0
	}
	return r.ByteOrder.Uint64(b)
}

// Byte decodes an uint64 from the Reader.
func (r *ReadChain) Byte(b *byte) *ReadChain {
	*b = byte(r.uint8())
	return r
}

// Bool decodes a bool from the Reader.
func (r *ReadChain) Bool(b *bool) *ReadChain {
	*b = r.uint8() != 0
	return r
}

// not read (present) because read (past participle) is in the struct's
// fields.
func (r *ReadChain) rd(amt int) []byte {
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
