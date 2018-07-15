[![GoDoc](https://godoc.org/howl.moe/binary?status.svg)](https://godoc.org/howl.moe/binary) [![Build Status](https://travis-ci.org/thehowl/binary.svg?branch=master)](https://travis-ci.org/thehowl/binary)

# binary

A faster binary encoder.

```
go get howl.moe/binary
```

## Migrating to version 2

All slice-related methods have been removed - this is because they allocated
their own slices (for reads). The way slices are implemented in binary protocols
is often arbitrary, some specify the length, some only in certain cases, some
use varints, others use different widths for the length. For this reason, we
removed the methods.

ByteSlice and String have been kept, since those are even used internally.

There is no way to read a string directly like there was in the previous
version - this is because of the aforementioned variability in implementation
about encoding the length. You can still use the newly-added `Read` method
(implements `io.Reader`) passing a byte slice with the length desired.

### Why you should

The biggest, yet simplest, change was by removing all slice allocations. This is
a common trick used often, in places such as
[fasthttp](https://github.com/valyala/fasthttp) and
[nanojson](https://howl.moe/nanojson) (yes, that's what we call shameless
advertising.)

Previously, every single tiny read and write allocated a byte slice. This is
actually quite expensive - it is a heap allocation, which needs to be tracked
by the garbage collector, and so on. Now the byte slice is pre-allocated in the
Reader/WriteChain itself (as an array, `buf [8]byte`, which is then used as a
slice, ie. `c.buf[:]`). This gives significant performance boosts:

```
$ git checkout v1
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/thehowl/binary
BenchmarkWriteSmall-4    	30000000	        36.4 ns/op
BenchmarkWriteMedium-4   	 5000000	       263 ns/op
BenchmarkWriteLong-4     	 1000000	      1649 ns/op
PASS
ok  	github.com/thehowl/binary	4.421s

$ git checkout master
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/thehowl/binary
BenchmarkWriteSmall-4    	100000000	        20.0 ns/op
BenchmarkWriteMedium-4   	20000000	        91.7 ns/op
BenchmarkWriteLong-4     	 3000000	       412 ns/op
PASS
ok  	github.com/thehowl/binary	5.630s
```

As you can see, for writes of large chunks of data, this can be up to a 4x
improvement.
