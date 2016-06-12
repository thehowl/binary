package binary

// Int8 encodes an int8 into the writer.
func (c *WriteChain) Int8(i int8) *WriteChain {
	return c.Uint8(uint8(i))
}

// Int16 encodes an int16 into the writer.
func (c *WriteChain) Int16(i int16) *WriteChain {
	return c.Uint16(uint16(i))
}

// Int32 encodes an int32 into the writer.
func (c *WriteChain) Int32(i int32) *WriteChain {
	return c.Uint32(uint32(i))
}

// Int64 encodes an int64 into the writer.
func (c *WriteChain) Int64(i int64) *WriteChain {
	return c.Uint64(uint64(i))
}

// Rune encodes a rune into the writer.
func (c *WriteChain) Rune(r rune) *WriteChain {
	return c.Int32(int32(r))
}
