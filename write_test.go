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
		BoolSlice([]bool{true, false, true}).
		Complex64(complex64(89723.3432)).
		Complex64Slice([]complex64{125, 145, 12, 51}).
		Complex128(complex128(24515.918681)).
		Complex128Slice([]complex128{1956156, 1, 412, 4, 124, 12, 41, 41, 4}).
		Float64(3.14).
		Float64Slice([]float64{3.14, 41, 1337.7331}).
		Float32(14.3).
		Float32Slice([]float32{15, 1, 512, 51251.1451, 1251.515}).
		Int8(-128).
		Int8Slice([]int8{121, 45, 42}).
		Int16(5195).
		Int16Slice([]int16{32767, 1337, -1, -32768}).
		Rune(912869241). // int32 == rune, so let's just test runes
		RuneSlice([]rune("akojgakgjlkejjyl")).
		Int32Slice([]int32{8968, 24523, 5353}).
		Int64(-9510515010901591).
		Int64Slice([]int64{1569018691, -125681968129362, 6236, 235623, -72, 723, 73, 57, 347, 73, 57, 357}).
		Byte(255). // byte == uint8, so let's just test bytes
		ByteSlice([]byte("世界")).
		Uint8Slice([]uint8{255, 254, 245, 244, 233}).
		Uint16(18956).
		Uint16Slice([]uint16{3453, 23442, 3}).
		Uint32(91681296).
		Uint32Slice([]uint32{}).
		Uint32Slice([]uint32{2501951, 2542, 34535}).
		Uint64(96891268928692).
		Uint64Slice([]uint64{9286926, 26, 26, 23, 5623, 623, 6, 236, 23623623}).
		String("We're finally done.").
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
			Uint64Slice([]uint64{52, 2342, 234, 19869186912682}).
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
			String("Ok meme").
			Int8(-122).
			Uint64Slice([]uint64{52, 2342, 234, 19869186912682}).
			Int32(186928).
			RuneSlice([]rune("how much wood would a woodchuck chuck if a woodchuck could chhuck wood?")).
			Int64Slice([]int64{6928692348, 242623, 235234523}).
			End()
	}
}

func Example_Write() {
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
		fmt.Println(err)
		return
	}
	fmt.Printf("% x\n", buf.Bytes())
	// Output: 0a 01 01 02 00 00 00
}
