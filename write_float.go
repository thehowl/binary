package binary

import (
	"math"
)

// Float32 encodes a float32 into the writer.
func (c *WriteChain) Float32(f float32) *WriteChain {
	return c.Uint32(math.Float32bits(f))
}

// Float64 encodes a float64 into the writer.
func (c *WriteChain) Float64(f float64) *WriteChain {
	return c.Uint64(math.Float64bits(f))
}

// Complex64 encodes a complex64 into the writer.
func (c *WriteChain) Complex64(com complex64) *WriteChain {
	return c.
		Float32(real(com)).
		Float32(imag(com))
}

// Complex128 encodes a complex128 into the writer.
func (c *WriteChain) Complex128(com complex128) *WriteChain {
	return c.
		Float64(real(com)).
		Float64(imag(com))
}
