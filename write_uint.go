package binary

import (
	"io"
	"sync"
)

var bufpool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 8)
	},
}

func (c *WriteChain) buf() []byte {
	if c._buf == nil {
		c._buf = bufpool.Get().([]byte)
	}
	return c._buf
}

// Uint8 encodes an uint8 into the writer.
func (c *WriteChain) Uint8(u uint8) *WriteChain {
	b := c.buf()
	b[0] = byte(u)
	return c.Bytes(b[:1])
}

// Uint16 encodes an uint16 into the writer.
func (c *WriteChain) Uint16(u uint16) *WriteChain {
	b := c.buf()
	c.ByteOrder.PutUint16(b[:2], u)
	return c.Bytes(b[:2])
}

// Uint32 encodes an uint32 into the writer.
func (c *WriteChain) Uint32(u uint32) *WriteChain {
	b := c.buf()
	c.ByteOrder.PutUint32(b[:4], u)
	return c.Bytes(b[:4])
}

// Uint64 encodes an uint64 into the writer.
func (c *WriteChain) Uint64(u uint64) *WriteChain {
	b := c.buf()
	c.ByteOrder.PutUint64(b, u)
	return c.Bytes(b)
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

// Bytes is a wrapper around c.Write which discards the error.
func (c *WriteChain) Bytes(b []byte) *WriteChain {
	i, err := c.Write(b)
	if err == nil && i != len(b) {
		c.err = io.ErrShortWrite
	}
	return c
}
