package binary_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/thehowl/binary"
)

const unmsh = "\x0b\x00\x10\x01\xff"

func TestRead(t *testing.T) {
	buf := bytes.NewBuffer(
		[]byte(unmsh),
	)
	r := &binary.Reader{
		Reader:    buf,
		ByteOrder: binary.LittleEndian,
	}
	{
		x := r.Int16()
		if x != 11 {
			t.Error("meme!", x)
		}
	}
	{
		x := r.Byte()
		if x != 16 {
			t.Error("meme!", x)
		}
	}
	{
		x := r.Bool()
		if !x {
			t.Error("meme!", x)
		}
	}
	{
		x := r.Int8()
		if x != -1 {
			t.Error("meme!", x)
		}
	}
	// more tests soon™
}

func ExampleReader() {
	buf := bytes.NewBuffer([]byte("\x05\xff\x02\x00"))
	r := &binary.Reader{
		Reader:    buf,
		ByteOrder: binary.LittleEndian,
	}
	var (
		b1 = r.Byte()
		b2 = r.Int8()
		i  = r.Int16()
	)
	_, err := r.End()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d - %d - %d\n", b1, b2, i)
	// Output: 5 - -1 - 2
}
