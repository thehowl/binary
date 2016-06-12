package binary

// Uint8 encodes an uint8 into the writer.
func (c *WriteChain) Uint8(u uint8) *WriteChain {
	return c.bytes([]byte{byte(u)})
}

// Uint16 encodes an uint16 into the writer.
func (c *WriteChain) Uint16(u uint16) *WriteChain {
	x := make([]byte, 2)
	c.ByteOrder.PutUint16(x, u)
	return c.bytes(x)
}

// Uint32 encodes an uint32 into the writer.
func (c *WriteChain) Uint32(u uint32) *WriteChain {
	x := make([]byte, 4)
	c.ByteOrder.PutUint32(x, u)
	return c.bytes(x)
}

// Uint64 encodes an uint64 into the writer.
func (c *WriteChain) Uint64(u uint64) *WriteChain {
	x := make([]byte, 8)
	c.ByteOrder.PutUint64(x, u)
	return c.bytes(x)
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

func (c *WriteChain) bytes(b []byte) *WriteChain {
	if c == nil || c.err != nil {
		return c
	}
	written, err := c.Writer.Write(b)
	c.written += written
	c.err = err
	if err == nil {
		c.err = checkWritten(written, len(b))
	}
	return c
}
