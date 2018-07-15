package binary

import (
	"io"
)

// Uint8 encodes an uint8 into the writer.
func (c *WriteChain) Uint8(u uint8) *WriteChain {
	c.buf[0] = byte(u)
	return c.ByteSlice(c.buf[:1])
}

// Uint16 encodes an uint16 into the writer.
func (c *WriteChain) Uint16(u uint16) *WriteChain {
	c.ByteOrder.PutUint16(c.buf[:2], u)
	return c.ByteSlice(c.buf[:2])
}

// Uint32 encodes an uint32 into the writer.
func (c *WriteChain) Uint32(u uint32) *WriteChain {
	c.ByteOrder.PutUint32(c.buf[:4], u)
	return c.ByteSlice(c.buf[:4])
}

// Uint64 encodes an uint64 into the writer.
func (c *WriteChain) Uint64(u uint64) *WriteChain {
	c.ByteOrder.PutUint64(c.buf[:], u)
	return c.ByteSlice(c.buf[:])
}

// Byte encodes a byte into the writer.
func (c *WriteChain) Byte(b byte) *WriteChain {
	return c.Uint8(uint8(b))
}

// Bool encodes a boolean into the write, transforming true into uint8(1) and
// false into uint8(0).
func (c *WriteChain) Bool(b bool) *WriteChain {
	if b {
		return c.Uint8(1)
	}
	return c.Uint8(0)
}

// ByteSlice writes a raw byte slice into the WriteChain
func (c *WriteChain) ByteSlice(b []byte) *WriteChain {
	i, err := c.Write(b)
	if err == nil && i != len(b) {
		c.err = io.ErrShortWrite
	}
	return c
}

// String writes a string to the write chain.
func (c *WriteChain) String(s string) *WriteChain {
	return c.ByteSlice([]byte(s))
}
