package binary

// BoolSlice encodes a slice of bools into the writer.
func (c *WriteChain) BoolSlice(b []bool) *WriteChain {
	for _, el := range b {
		c.Bool(el)
	}
	return c
}

// ByteSlice encodes a slice of bytes into the writer.
func (c *WriteChain) ByteSlice(b []byte) *WriteChain {
	for _, el := range b {
		c.Byte(el)
	}
	return c
}

// Complex128Slice encodes a slice of complex128s into the writer.
func (c *WriteChain) Complex128Slice(com []complex128) *WriteChain {
	for _, el := range com {
		c.Complex128(el)
	}
	return c
}

// Complex64Slice encodes a slice of complex64s into the writer.
func (c *WriteChain) Complex64Slice(com []complex64) *WriteChain {
	for _, el := range com {
		c.Complex64(el)
	}
	return c
}

// Float32Slice encodes a slice of float32s into the writer.
func (c *WriteChain) Float32Slice(f []float32) *WriteChain {
	for _, el := range f {
		c.Float32(el)
	}
	return c
}

// Float64Slice encodes a slice of float64s into the writer.
func (c *WriteChain) Float64Slice(f []float64) *WriteChain {
	for _, el := range f {
		c.Float64(el)
	}
	return c
}

// Int16Slice encodes a slice of int16s into the writer.
func (c *WriteChain) Int16Slice(i []int16) *WriteChain {
	for _, el := range i {
		c.Int16(el)
	}
	return c
}

// Int32Slice encodes a slice of int32s into the writer.
func (c *WriteChain) Int32Slice(i []int32) *WriteChain {
	for _, el := range i {
		c.Int32(el)
	}
	return c
}

// Int64Slice encodes a slice of int64s into the writer.
func (c *WriteChain) Int64Slice(i []int64) *WriteChain {
	for _, el := range i {
		c.Int64(el)
	}
	return c
}

// Int8Slice encodes a slice of int8s into the writer.
func (c *WriteChain) Int8Slice(i []int8) *WriteChain {
	for _, el := range i {
		c.Int8(el)
	}
	return c
}

// RuneSlice encodes a slice of runes into the writer.
func (c *WriteChain) RuneSlice(r []rune) *WriteChain {
	for _, el := range r {
		c.Rune(el)
	}
	return c
}

// Uint16Slice encodes a slice of uint16s into the writer.
func (c *WriteChain) Uint16Slice(u []uint16) *WriteChain {
	for _, el := range u {
		c.Uint16(el)
	}
	return c
}

// Uint32Slice encodes a slice of uint32s into the writer.
func (c *WriteChain) Uint32Slice(u []uint32) *WriteChain {
	for _, el := range u {
		c.Uint32(el)
	}
	return c
}

// Uint64Slice encodes a slice of uint64s into the writer.
func (c *WriteChain) Uint64Slice(u []uint64) *WriteChain {
	for _, el := range u {
		c.Uint64(el)
	}
	return c
}

// Uint8Slice encodes a slice of uint8s into the writer.
func (c *WriteChain) Uint8Slice(u []uint8) *WriteChain {
	for _, el := range u {
		c.Uint8(el)
	}
	return c
}

// String encodes a string into the writer.
func (c *WriteChain) String(s string) *WriteChain {
	return c.ByteSlice([]byte(s))
}
