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
	r := &binary.ReadChain{
		Reader:    buf,
		ByteOrder: binary.LittleEndian,
	}
	{
		var x int16
		r.Int16(&x)
		if x != 11 {
			t.Error("meme!", x)
		}
	}
	{
		var x byte
		r.Byte(&x)
		if x != 16 {
			t.Error("meme!", x)
		}
	}
	{
		var x bool
		r.Bool(&x)
		if !x {
			t.Error("meme!", x)
		}
	}
	{
		var x int8
		r.Int8(&x)
		if x != -1 {
			t.Error("meme!", x)
		}
	}
	// more tests soon™
}

func ExampleReadChain() {
	buf := bytes.NewBuffer([]byte("\x05\xff\x02\x00"))
	r := &binary.ReadChain{
		Reader:    buf,
		ByteOrder: binary.LittleEndian,
	}
	var (
		b1 byte
		b2 int8
		i  int16
	)
	_, err := r.
		Byte(&b1).
		Int8(&b2).
		Int16(&i).
		End()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d - %d - %d\n", b1, b2, i)
	// Output: 5 - -1 - 2
}
