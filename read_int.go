package binary

// Int8 decodes an int8 from the Reader.
func (r *ReadChain) Int8(i *int8) *ReadChain {
	*i = int8(r.uint8())
	return r
}

// Int16 decodes an int16 from the Reader.
func (r *ReadChain) Int16(i *int16) *ReadChain {
	*i = int16(r.uint16())
	return r
}

// Int32 decodes an int32 from the Reader.
func (r *ReadChain) Int32(i *int32) *ReadChain {
	*i = int32(r.uint32())
	return r
}

// Int64 decodes an int64 from the Reader.
func (r *ReadChain) Int64(i *int64) *ReadChain {
	*i = int64(r.uint64())
	return r
}

// Rune decodes an rune from the Reader.
func (r *ReadChain) Rune(ru *rune) *ReadChain {
	*ru = rune(int32(r.uint32()))
	return r
}
