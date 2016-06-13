package binary

// Int8 decodes an int8 from the Reader.
func (r *Reader) Int8() int8 {
	return int8(r.Uint8())
}

// Int16 decodes an int16 from the Reader.
func (r *Reader) Int16() int16 {
	return int16(r.Uint16())
}

// Int32 decodes an int32 from the Reader.
func (r *Reader) Int32() int32 {
	return int32(r.Uint32())
}

// Int64 decodes an int64 from the Reader.
func (r *Reader) Int64() int64 {
	return int64(r.Uint64())
}

// Rune decodes an rune from the Reader.
func (r *Reader) Rune() rune {
	return rune(int32(r.Uint32()))
}
