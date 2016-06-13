package binary

// BoolSlice decodes a slice of bools.
func (r *Reader) BoolSlice(length int) []bool {
	sl := make([]bool, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Bool()
	}
	return sl
}

// ByteSlice decodes a slice of bytes.
func (r *Reader) ByteSlice(length int) []byte {
	return r.rd(length)
}

// Int16Slice decodes a slice of int16s.
func (r *Reader) Int16Slice(length int) []int16 {
	sl := make([]int16, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Int16()
	}
	return sl
}

// Int32Slice decodes a slice of int32s.
func (r *Reader) Int32Slice(length int) []int32 {
	sl := make([]int32, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Int32()
	}
	return sl
}

// Int64Slice decodes a slice of int64s.
func (r *Reader) Int64Slice(length int) []int64 {
	sl := make([]int64, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Int64()
	}
	return sl
}

// Int8Slice decodes a slice of int8s.
func (r *Reader) Int8Slice(length int) []int8 {
	sl := make([]int8, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Int8()
	}
	return sl
}

// RuneSlice decodes a slice of runes.
func (r *Reader) RuneSlice(length int) []rune {
	return []rune(string(r.rd(length)))
}

// Uint16Slice decodes a slice of uint16s.
func (r *Reader) Uint16Slice(length int) []uint16 {
	sl := make([]uint16, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Uint16()
	}
	return sl
}

// Uint32Slice decodes a slice of uint32s.
func (r *Reader) Uint32Slice(length int) []uint32 {
	sl := make([]uint32, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Uint32()
	}
	return sl
}

// Uint64Slice decodes a slice of uint64s.
func (r *Reader) Uint64Slice(length int) []uint64 {
	sl := make([]uint64, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Uint64()
	}
	return sl
}

// Uint8Slice decodes a slice of uint8s from the Reader.
func (r *Reader) Uint8Slice(length int) []uint8 {
	sl := make([]uint8, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Uint8()
	}
	return sl
}

// Float32Slice decodes a slice of float32s from the Reader.
func (r *Reader) Float32Slice(length int) []float32 {
	sl := make([]float32, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Float32()
	}
	return sl
}

// Float64Slice decodes a slice of float64s from the Reader.
func (r *Reader) Float64Slice(length int) []float64 {
	sl := make([]float64, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Float64()
	}
	return sl
}

// Complex64Slice decodes a slice of complex64s from the Reader.
func (r *Reader) Complex64Slice(length int) []complex64 {
	sl := make([]complex64, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Complex64()
	}
	return sl
}

// Complex128Slice decodes a slice of complex128s from the Reader.
func (r *Reader) Complex128Slice(length int) []complex128 {
	sl := make([]complex128, length)
	for i := 0; i < length; i++ {
		sl[i] = r.Complex128()
	}
	return sl
}

// String decodes a string from the Reader.
func (r *Reader) String(length int) string {
	return string(r.rd(length))
}
