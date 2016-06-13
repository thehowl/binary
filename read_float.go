package binary

import (
	"math"
)

// Float32 decodes a float32 from the Reader.
func (r *Reader) Float32() float32 {
	return math.Float32frombits(r.Uint32())
}

// Float64 decodes a float64 from the Reader.
func (r *Reader) Float64() float64 {
	return math.Float64frombits(r.Uint64())
}

// Complex64 decodes a Complex64 from the Reader.
func (r *Reader) Complex64() complex64 {
	return complex(r.Float32(), r.Float32())
}

// Complex128 decodes a Complex128 from the Reader.
func (r *Reader) Complex128() complex128 {
	return complex(r.Float64(), r.Float64())
}
