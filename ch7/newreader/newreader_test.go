package main

import (
	"testing"
)

func TestReadInSmallPieces(t *testing.T) {
	r := NewReader("Hello, World!")

	frags := []struct {
		n     int
		bs    []byte
		exp   string
		ioerr bool
	}{
		{n: 3,
			bs:  make([]byte, 3),
			exp: "Hel",
		},
		{n: 10,
			bs:  make([]byte, 11),
			exp: "lo, World!\x00",
		},
		{
			ioerr: true,
		},
	}

	for _, frag := range frags {
		n, err := r.Read(frag.bs)
		switch {
		case n != frag.n:
			t.Errorf("wrong value n; got: %v; exp: %v", n, frag.n)
		case err != nil && !frag.ioerr:
			t.Errorf("got unexpected error: %v", err)
		case err == nil && frag.ioerr:
			t.Errorf("expected error, got nil")
		case string(frag.bs) != frag.exp:
			t.Errorf("wrong value err; got: %v; exp: %v", frag.bs, frag.exp)
		}
	}
}
