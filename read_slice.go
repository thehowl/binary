package binary

// BoolSlice decodes a slice of bools.
func (r *ReadChain) BoolSlice(b []bool) *ReadChain {
	for index, el := range b {
		r.Bool(&el)
		b[index] = el
	}
	return r
}

// ByteSlice decodes a slice of bytes.
func (r *ReadChain) ByteSlice(b []byte) *ReadChain {
	x := r.rd(len(b))
	copy(b, x)
	return r
}

// Int16Slice decodes a slice of int16s.
func (r *ReadChain) Int16Slice(i []int16) *ReadChain {
	for index, el := range i {
		r.Int16(&el)
		i[index] = el
	}
	return r
}

// Int32Slice decodes a slice of int32s.
func (r *ReadChain) Int32Slice(i []int32) *ReadChain {
	for index, el := range i {
		r.Int32(&el)
		i[index] = el
	}
	return r
}

// Int64Slice decodes a slice of int64s.
func (r *ReadChain) Int64Slice(i []int64) *ReadChain {
	for index, el := range i {
		r.Int64(&el)
		i[index] = el
	}
	return r
}

// Int8Slice decodes a slice of int8s.
func (r *ReadChain) Int8Slice(i []int8) *ReadChain {
	for index, el := range i {
		r.Int8(&el)
		i[index] = el
	}
	return r
}

// RuneSlice decodes a slice of runes.
func (r *ReadChain) RuneSlice(ru []rune) *ReadChain {
	lb := []byte(string(ru))
	r.ByteSlice(lb)
	copy(ru, []rune(string(lb)))
	return r
}

// Uint16Slice decodes a slice of uint16s.
func (r *ReadChain) Uint16Slice(u []uint16) *ReadChain {
	for index, el := range u {
		r.Uint16(&el)
		u[index] = el
	}
	return r
}

// Uint32Slice decodes a slice of uint32s.
func (r *ReadChain) Uint32Slice(u []uint32) *ReadChain {
	for index, el := range u {
		r.Uint32(&el)
		u[index] = el
	}
	return r
}

// Uint64Slice decodes a slice of uint64s.
func (r *ReadChain) Uint64Slice(u []uint64) *ReadChain {
	for index, el := range u {
		r.Uint64(&el)
		u[index] = el
	}
	return r
}

// Uint8Slice decodes a slice of uint8s from the Reader.
func (r *ReadChain) Uint8Slice(u []uint8) *ReadChain {
	for index, el := range u {
		r.Uint8(&el)
		u[index] = el
	}
	return r
}

// Float32Slice decodes a slice of float32s from the Reader.
func (r *ReadChain) Float32Slice(u []float32) *ReadChain {
	for index := range u {
		u[index] = r.float32()
	}
	return r
}

// Float64Slice decodes a slice of float64s from the Reader.
func (r *ReadChain) Float64Slice(u []float64) *ReadChain {
	for index := range u {
		u[index] = r.float64()
	}
	return r
}

// Complex64Slice decodes a slice of complex64s from the Reader.
func (r *ReadChain) Complex64Slice(u []complex64) *ReadChain {
	for index, el := range u {
		r.Complex64(&el)
		u[index] = el
	}
	return r
}

// Complex128Slice decodes a slice of complex128s from the Reader.
func (r *ReadChain) Complex128Slice(u []complex128) *ReadChain {
	for index, el := range u {
		r.Complex128(&el)
		u[index] = el
	}
	return r
}

// String decodes a string from the Reader. l is the desired length of the
// string, as in the desired amount of bytes that should be read from the
// Reader.
func (r *ReadChain) String(s *string, l int) *ReadChain {
	into := make([]byte, l)
	r.ByteSlice(into)
	*s = string(into)
	return r
}
