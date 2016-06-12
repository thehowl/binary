package binary

import (
	"math"
)

// Float32 decodes a float32 from the Reader.
func (r *ReadChain) Float32(f *float32) *ReadChain {
	*f = r.float32()
	return r
}

func (r *ReadChain) float32() float32 {
	return math.Float32frombits(r.uint32())
}

// Float64 decodes a float64 from the Reader.
func (r *ReadChain) Float64(f *float64) *ReadChain {
	*f = r.float64()
	return r
}

func (r *ReadChain) float64() float64 {
	return math.Float64frombits(r.uint64())
}

// Complex64 decodes a Complex64 from the Reader.
func (r *ReadChain) Complex64(c *complex64) *ReadChain {
	*c = complex(r.float32(), r.float32())
	return r
}

// Complex128 decodes a Complex128 from the Reader.
func (r *ReadChain) Complex128(c *complex128) *ReadChain {
	*c = complex(r.float64(), r.float64())
	return r
}
