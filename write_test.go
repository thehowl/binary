package binary_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/thehowl/binary"
)

func TestWrite(t *testing.T) {
	buf := &bytes.Buffer{}
	w := &binary.WriteChain{
		Writer:    buf,
		ByteOrder: binary.LittleEndian,
	}
	written, err := w.
		Bool(true).
		Complex64(complex64(89723.3432)).
		Complex128(complex128(24515.918681)).
		Float64(3.14).
		Float32(14.3).
		Int8(-128).
		Int16(5195).
		Rune(912869241). // int32 == rune, so let's just test runes
		Int64(-9510515010901591).
		Byte(255). // byte == uint8, so let's just test bytes
		Uint16(18956).
		Uint32(91681296).
		Uint64(96891268928692).
		End()
	t.Log(written, "bytes written")
	// t.Logf("%x", buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}
}

type fakeWriter struct{}

func (f fakeWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func BenchmarkWriteSmall(b *testing.B) {
	var x = &binary.WriteChain{
		Writer:    &fakeWriter{},
		ByteOrder: binary.LittleEndian,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.
			Byte(255).
			End()
	}
}

func BenchmarkWriteMedium(b *testing.B) {
	var x = &binary.WriteChain{
		Writer:    &fakeWriter{},
		ByteOrder: binary.LittleEndian,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.
			Int8(-122).
			Uint64(52).
			Uint64(2342).
			Uint64(234).
			Uint64(19869186912682).
			End()
	}
}

func BenchmarkWriteLong(b *testing.B) {
	var x = &binary.WriteChain{
		Writer:    &fakeWriter{},
		ByteOrder: binary.LittleEndian,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.
			ByteSlice([]byte("Ok meme")).
			Int8(-122).
			Uint64(52).
			Uint64(2342).
			Uint64(234).
			Uint64(19869186912682).
			Int32(186928).
			ByteSlice([]byte("how much wood would a woodchuck chuck if a woodchuck could chhuck wood?")).
			Int64(6928692348).
			Int64(242623).
			Int64(235234523).
			End()
	}
}

func ExampleWriteChain() {
	buf := &bytes.Buffer{}
	writer := &binary.WriteChain{
		Writer:    buf,
		ByteOrder: binary.LittleEndian,
	}
	_, err := writer.
		Uint16(266).
		Byte(1).
		Uint32(2).
		End()
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x\n", buf.Bytes())
	// Output: 0a 01 01 02 00 00 00
}
